// Package pipevisitor converts raw .pipe content into a Pipe using the ANTLR4-generated grammar.
// Run scripts/generate-pipe-grammar.sh to regenerate grammar/ after editing the .g4 files.
package pipevisitor

import (
	"errors"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	grammar "github.com/gear6io/pragmata/pkg/grammars/pipesgrammar"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/gear6io/pragmata/pkg/valuer"
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

func (v *pipeVisitor) VisitTypeDir(ctx *grammar.TypeDirContext) interface{} {
	v.pipe.Type = mapPipeType(ctx.VALUE().GetText())
	return nil
}

func (v *pipeVisitor) VisitNameDir(ctx *grammar.NameDirContext) interface{} {
	v.pipe.Name = ctx.VALUE().GetText()
	return nil
}

func (v *pipeVisitor) VisitDescriptionDir(ctx *grammar.DescriptionDirContext) interface{} {
	v.pipe.Description = ctx.VALUE().GetText()
	return nil
}

func (v *pipeVisitor) VisitDescriptionMLDir(ctx *grammar.DescriptionMLDirContext) interface{} {
	lines := make([]string, 0, len(ctx.AllSECTION_LINE()))
	for _, tok := range ctx.AllSECTION_LINE() {
		if s := strings.TrimSpace(tok.GetText()); s != "" {
			lines = append(lines, s)
		}
	}
	v.pipe.Description = strings.Join(lines, " ")
	return nil
}

func (v *pipeVisitor) VisitTagsDir(ctx *grammar.TagsDirContext) interface{} {
	raw := strings.Trim(ctx.VALUE().GetText(), "[] \t")
	for _, tag := range strings.Split(raw, ",") {
		if t := strings.TrimSpace(tag); t != "" {
			v.pipe.Tags = append(v.pipe.Tags, t)
		}
	}
	return nil
}

func (v *pipeVisitor) VisitOwnerDir(ctx *grammar.OwnerDirContext) interface{} {
	v.pipe.Owner = ctx.VALUE().GetText()
	return nil
}

func (v *pipeVisitor) VisitDestinationDir(ctx *grammar.DestinationDirContext) interface{} {
	v.pipe.Destination = ctx.VALUE().GetText()
	return nil
}

func (v *pipeVisitor) VisitScheduleDir(ctx *grammar.ScheduleDirContext) interface{} {
	v.pipe.Schedule = ctx.VALUE().GetText()
	return nil
}

func (v *pipeVisitor) VisitSourcesClause(ctx *grammar.SourcesClauseContext) interface{} {
	for _, e := range ctx.AllSource() {
		e.(antlr.ParseTree).Accept(v)
	}
	return nil
}

func (v *pipeVisitor) VisitAliasedSource(ctx *grammar.AliasedSourceContext) interface{} {
	v.pipe.Sources = append(v.pipe.Sources, pipetypes.Source{
		Alias: ctx.GetAlias().GetText(),
		Table: ctx.GetTable().GetText(),
	})
	return nil
}

func (v *pipeVisitor) VisitSimpleSource(ctx *grammar.SimpleSourceContext) interface{} {
	name := ctx.GetName().GetText()
	v.pipe.Sources = append(v.pipe.Sources, pipetypes.Source{Alias: name, Table: name})
	return nil
}

func (v *pipeVisitor) VisitParamsClause(ctx *grammar.ParamsClauseContext) interface{} {
	for _, e := range ctx.AllParam() {
		e.(antlr.ParseTree).Accept(v)
	}
	return nil
}

func (v *pipeVisitor) VisitParam(ctx *grammar.ParamContext) interface{} {
	name := strings.TrimSpace(ctx.GetPname().GetText())
	dtype := ctx.GetDtype().GetText()
	var defVal string
	if pv := ctx.ParamValue(); pv != nil {
		defVal = strings.Trim(pv.GetText(), "\"")
	}
	v.pipe.Params = append(v.pipe.Params, pipetypes.ParamDef{
		Name: name, DataType: dtype, DefaultValue: defVal,
	})
	return nil
}

func (v *pipeVisitor) VisitPipelineClause(ctx *grammar.PipelineClauseContext) interface{} {
	var cur *pipetypes.Node
	for _, tok := range ctx.AllSECTION_LINE() {
		line := strings.TrimSpace(tok.GetText())
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "@") {
			if cur != nil {
				v.pipe.Nodes = append(v.pipe.Nodes, *cur)
			}
			header := strings.TrimPrefix(line, "@")
			if idx := strings.IndexByte(header, ':'); idx >= 0 {
				header = header[:idx]
			}
			name := strings.TrimSpace(header)
			if name == "" {
				v.errs = append(v.errs, fmt.Errorf("pipeline node has empty name"))
				return nil
			}
			cur = &pipetypes.Node{Name: name}
		} else if cur != nil {
			if cur.SQL != "" {
				cur.SQL += " "
			}
			cur.SQL += line
		}
	}
	if cur != nil {
		v.pipe.Nodes = append(v.pipe.Nodes, *cur)
	}
	return nil
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
