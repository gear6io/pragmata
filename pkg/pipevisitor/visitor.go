// Package pipevisitor visits .pipe file content using an ANTLR4-generated grammar.
// Run scripts/generate-pipe-grammar.sh to regenerate grammar/ after editing PipeLangLexer.g4 or PipeLang.g4.
package pipevisitor

import (
	"errors"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	grammar "github.com/gear6io/pragmata/pkg/grammars/pipesgrammar"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/gear6io/pragmata/pkg/valuer"
	"gopkg.in/yaml.v3"
)

// pipeVisitor implements grammar.PipeLangVisitor, accumulating a Pipe and any errors.
type pipeVisitor struct {
	grammar.BasePipeLangVisitor
	pipe *pipetypes.Pipe
	errs []error
}

// Visit converts raw .pipe file content into a Pipe. name is the initial pipe
// name (typically derived from the filename); a name: directive in the file
// overrides it.
func Visit(name, content string) (*pipetypes.Pipe, error) {
	lexer := grammar.NewPipeLangLexer(antlr.NewInputStream(content))
	lexErr := &errListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexErr)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := grammar.NewPipeLang(stream)
	parseErr := &errListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(parseErr)

	tree := p.PipeFile()

	if lexErr.msg != "" {
		return nil, fmt.Errorf("lex error in %q: %s", name, lexErr.msg)
	}
	if parseErr.msg != "" {
		return nil, fmt.Errorf("parse error in %q: %s", name, parseErr.msg)
	}

	v := &pipeVisitor{
		pipe: &pipetypes.Pipe{Name: name, Content: content},
	}
	tree.Accept(v)

	if len(v.errs) > 0 {
		return nil, errors.Join(v.errs...)
	}
	if len(v.pipe.Nodes) == 0 {
		return nil, fmt.Errorf("pipe %q has no pipeline nodes", name)
	}
	if v.pipe.Type == pipetypes.PipeTypeUndefined {
		return nil, fmt.Errorf("pipe %q has no type declaration", name)
	}
	return v.pipe, nil
}

// VisitPipeFile drives traversal so child Accept calls dispatch through *pipeVisitor.
func (v *pipeVisitor) VisitPipeFile(ctx *grammar.PipeFileContext) interface{} {
	for _, dir := range ctx.AllDirective() {
		dir.(antlr.ParseTree).Accept(v)
	}
	return nil
}

func (v *pipeVisitor) VisitType(ctx *grammar.TypeContext) interface{} {
	v.pipe.Type = mapPipeType(ctx.REST_OF_LINE().GetText())
	return nil
}

func (v *pipeVisitor) VisitName(ctx *grammar.NameContext) interface{} {
	v.pipe.Name = ctx.REST_OF_LINE().GetText()
	return nil
}

func (v *pipeVisitor) VisitDescription(ctx *grammar.DescriptionContext) interface{} {
	v.pipe.Description = ctx.REST_OF_LINE().GetText()
	return nil
}

func (v *pipeVisitor) VisitDescriptionML(ctx *grammar.DescriptionMLContext) interface{} {
	lines := make([]string, 0, len(ctx.AllBLOCK_LINE()))
	for _, tok := range ctx.AllBLOCK_LINE() {
		lines = append(lines, strings.TrimSpace(tok.GetText()))
	}
	v.pipe.Description = strings.Join(lines, " ")
	return nil
}

func (v *pipeVisitor) VisitTags(ctx *grammar.TagsContext) interface{} {
	raw := strings.Trim(ctx.REST_OF_LINE().GetText(), "[] \t")
	for _, tag := range strings.Split(raw, ",") {
		if t := strings.TrimSpace(tag); t != "" {
			v.pipe.Tags = append(v.pipe.Tags, t)
		}
	}
	return nil
}

func (v *pipeVisitor) VisitOwner(ctx *grammar.OwnerContext) interface{} {
	v.pipe.Owner = ctx.REST_OF_LINE().GetText()
	return nil
}

func (v *pipeVisitor) VisitDestination(ctx *grammar.DestinationContext) interface{} {
	v.pipe.Destination = ctx.REST_OF_LINE().GetText()
	return nil
}

func (v *pipeVisitor) VisitSchedule(ctx *grammar.ScheduleContext) interface{} {
	v.pipe.Schedule = ctx.REST_OF_LINE().GetText()
	return nil
}

func (v *pipeVisitor) VisitSourcesClause(ctx *grammar.SourcesClauseContext) interface{} {
	for _, tok := range ctx.AllBLOCK_LINE() {
		src, err := parseSourceLine(tok.GetText())
		if err != nil {
			v.errs = append(v.errs, err)
			return nil
		}
		v.pipe.Sources = append(v.pipe.Sources, src)
	}
	return nil
}

func (v *pipeVisitor) VisitParamsClause(ctx *grammar.ParamsClauseContext) interface{} {
	for _, tok := range ctx.AllBLOCK_LINE() {
		pd, err := parseParamLine(tok.GetText())
		if err != nil {
			v.errs = append(v.errs, err)
			return nil
		}
		v.pipe.Params = append(v.pipe.Params, pd)
	}
	return nil
}

func (v *pipeVisitor) VisitPipelineClause(ctx *grammar.PipelineClauseContext) interface{} {
	for _, nodeCtx := range ctx.PipelineBlock().AllPipelineNode() {
		header := strings.TrimSpace(nodeCtx.NODE_HEADER().GetText())
		header = strings.TrimPrefix(header, "@")
		if idx := strings.IndexByte(header, ':'); idx >= 0 {
			header = header[:idx]
		}
		name := strings.TrimSpace(header)
		if name == "" {
			v.errs = append(v.errs, fmt.Errorf("pipeline node has empty name"))
			return nil
		}
		lines := make([]string, 0, len(nodeCtx.AllPRQL_LINE()))
		for _, tok := range nodeCtx.AllPRQL_LINE() {
			lines = append(lines, strings.TrimSpace(tok.GetText()))
		}
		v.pipe.Nodes = append(v.pipe.Nodes, pipetypes.Node{
			Name: name,
			SQL:  strings.Join(lines, " "),
		})
	}
	return nil
}

// parseSourceLine parses one BLOCK_LINE from the sources: section.
//
//	"  - raw_events: events # comment\n"  →  Source{Alias:"raw_events", Table:"events"}
//	"  - users\n"                          →  Source{Alias:"users",      Table:"users"}
func parseSourceLine(raw string) (pipetypes.Source, error) {
	line := strings.TrimSpace(strings.TrimRight(raw, "\r\n"))
	if idx := strings.Index(line, " #"); idx >= 0 { line = line[:idx] }
	line = strings.TrimPrefix(line, "- ")
	line = strings.TrimSpace(line)
	if line == "" {
		return pipetypes.Source{}, fmt.Errorf("empty source line")
	}
	parts := strings.SplitN(line, ":", 2)
	if len(parts) == 2 {
		alias := strings.TrimSpace(parts[0])
		table := strings.TrimSpace(parts[1])
		if alias == "" || table == "" {
			return pipetypes.Source{}, fmt.Errorf("invalid source entry: %q", line)
		}
		return pipetypes.Source{Alias: alias, Table: table}, nil
	}
	table := strings.TrimSpace(parts[0])
	return pipetypes.Source{Alias: table, Table: table}, nil
}

// parseParamLine parses one BLOCK_LINE from the params: section.
//
//	"  country: { type: string, default: \"US\" }\n"
//	→ ParamDef{Name:"country", DataType:"string", DefaultValue:"US"}
func parseParamLine(raw string) (pipetypes.ParamDef, error) {
	line := strings.TrimSpace(strings.TrimRight(raw, "\r\n"))
	if idx := strings.Index(line, " #"); idx >= 0 { line = line[:idx] }
	if line == "" {
		return pipetypes.ParamDef{}, fmt.Errorf("empty param line")
	}
	var m map[string]struct {
		Type    string `yaml:"type"`
		Default string `yaml:"default"`
	}
	if err := yaml.Unmarshal([]byte(line), &m); err != nil || len(m) != 1 {
		return pipetypes.ParamDef{}, fmt.Errorf("invalid param entry: %q", line)
	}
	for name, v := range m {
		return pipetypes.ParamDef{Name: name, DataType: v.Type, DefaultValue: v.Default}, nil
	}
	return pipetypes.ParamDef{}, fmt.Errorf("invalid param entry: %q", line)
}

func mapPipeType(s string) pipetypes.PipeType {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "ENDPOINT":
		return pipetypes.PipeTypeEndpoint
	case "TABLE":
		return pipetypes.PipeTypeTable
	case "VIEW":
		return pipetypes.PipeTypeView
	case "INCREMENTAL":
		return pipetypes.PipeTypeIncremental
	case "SNAPSHOT":
		return pipetypes.PipeTypeSnapshot
	case "MATERIALIZED":
		return pipetypes.PipeTypeMaterialized
	case "COPY":
		return pipetypes.PipeTypeCopy
	default:
		return pipetypes.PipeType{String: valuer.NewString(strings.ToUpper(strings.TrimSpace(s)))}
	}
}


// errListener collects the first syntax error produced by the lexer or parser.
type errListener struct {
	*antlr.DefaultErrorListener
	msg string
}

func (e *errListener) SyntaxError(
	_ antlr.Recognizer,
	_ interface{},
	line, col int,
	msg string,
	_ antlr.RecognitionException,
) {
	if e.msg == "" {
		e.msg = fmt.Sprintf("line %d:%d %s", line, col, msg)
	}
}
