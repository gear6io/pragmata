// Package pipeparser parses .pipe file content using an ANTLR4-generated grammar.
// Run scripts/generate-pipe-grammar.sh to regenerate grammar/ after editing PipeLangLexer.g4 or PipeLang.g4.
package pipeparser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	grammar "github.com/gear6io/pragmata/pkg/grammars/pipesgrammar"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/gear6io/pragmata/pkg/valuer"
)

// Parse converts raw .pipe file content into a Pipe. name is the initial pipe
// name (typically derived from the filename); a name: directive in the file
// overrides it.
func Parse(name, content string) (*pipetypes.Pipe, error) {
	if !strings.HasSuffix(content, "\n") {
		content += "\n"
	}

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
	pipe := &pipetypes.Pipe{Name: name, Content: content}
	return pipe, parsePipe(tree, pipe)
}

func parsePipe(ctx grammar.IPipeFileContext, pipe *pipetypes.Pipe) error {
	var errs []error

	for _, dir := range ctx.AllDirective() {
		switch d := dir.(type) {
		case *grammar.TypeDirContext:
			pipe.Type = mapPipeType(stripComment(d.REST_OF_LINE().GetText()))
		case *grammar.NameDirContext:
			pipe.Name = stripComment(d.REST_OF_LINE().GetText())
		case *grammar.DescriptionDirContext:
			pipe.Description = stripComment(d.REST_OF_LINE().GetText())
		case *grammar.DescriptionMLDirContext:
			lines := collectBlockLines(d.AllBLOCK_LINE())
			pipe.Description = strings.TrimSpace(dedent(strings.Join(lines, "\n")))
		case *grammar.TagsDirContext:
			raw := strings.Trim(stripComment(d.REST_OF_LINE().GetText()), "[] \t")
			for _, tag := range strings.Split(raw, ",") {
				if t := strings.TrimSpace(tag); t != "" {
					pipe.Tags = append(pipe.Tags, t)
				}
			}
		case *grammar.OwnerDirContext:
			pipe.Owner = stripComment(d.REST_OF_LINE().GetText())
		case *grammar.DestinationDirContext:
			pipe.Destination = stripComment(d.REST_OF_LINE().GetText())
		case *grammar.ScheduleDirContext:
			pipe.Schedule = stripComment(d.REST_OF_LINE().GetText())
		case *grammar.SourcesDirContext:
			for _, tok := range d.AllBLOCK_LINE() {
				src, err := parseSourceLine(tok.GetText())
				if err != nil {
					errs = append(errs, err)
					break
				}
				pipe.Sources = append(pipe.Sources, src)
			}
		case *grammar.ParamsDirContext:
			for _, tok := range d.AllBLOCK_LINE() {
				pd, err := parseParamLine(tok.GetText())
				if err != nil {
					errs = append(errs, err)
					break
				}
				pipe.Params = append(pipe.Params, pd)
			}
		case *grammar.PipelineDirContext:
			nodes, err := parsePipelineBlock(d.PipelineBlock())
			if err != nil {
				errs = append(errs, err)
			} else {
				pipe.Nodes = nodes
			}
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return validate(pipe)
}

func parsePipelineBlock(ctx grammar.IPipelineBlockContext) ([]pipetypes.Node, error) {
	var nodes []pipetypes.Node
	for _, nodeCtx := range ctx.AllPipelineNode() {
		name := extractNodeName(nodeCtx.NODE_HEADER().GetText())
		if name == "" {
			return nil, fmt.Errorf("pipeline node has empty name")
		}
		lines := make([]string, 0, len(nodeCtx.AllPRQL_LINE()))
		for _, tok := range nodeCtx.AllPRQL_LINE() {
			lines = append(lines, strings.TrimRight(tok.GetText(), "\r\n"))
		}
		nodes = append(nodes, pipetypes.Node{
			Name: name,
			SQL:  strings.TrimSpace(dedent(strings.Join(lines, "\n"))),
		})
	}
	return nodes, nil
}

func validate(p *pipetypes.Pipe) error {
	if len(p.Nodes) == 0 {
		return fmt.Errorf("pipe %q has no pipeline nodes", p.Name)
	}
	if p.Type == pipetypes.PipeTypeUndefined {
		return fmt.Errorf("pipe %q has no type declaration", p.Name)
	}
	return nil
}

// ── String helpers ────────────────────────────────────────────────────────────

func stripComment(s string) string {
	s = strings.TrimRight(s, "\r\n")
	if idx := strings.Index(s, " #"); idx >= 0 {
		s = s[:idx]
	}
	return strings.TrimSpace(s)
}

func collectBlockLines(tokens []antlr.TerminalNode) []string {
	lines := make([]string, 0, len(tokens))
	for _, tok := range tokens {
		lines = append(lines, strings.TrimRight(tok.GetText(), "\r\n"))
	}
	return lines
}

// parseSourceLine parses one BLOCK_LINE from the sources: section.
//
//	"  - raw_events: events # comment\n"  →  Source{Alias:"raw_events", Table:"events"}
//	"  - users\n"                          →  Source{Alias:"users",      Table:"users"}
func parseSourceLine(raw string) (pipetypes.Source, error) {
	line := strings.TrimSpace(strings.TrimRight(raw, "\r\n"))
	line = stripComment(line)
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
	line = stripComment(line)
	if line == "" {
		return pipetypes.ParamDef{}, fmt.Errorf("empty param line")
	}
	colonIdx := strings.IndexByte(line, ':')
	if colonIdx < 0 {
		return pipetypes.ParamDef{}, fmt.Errorf("invalid param entry (no colon): %q", line)
	}
	name := strings.TrimSpace(line[:colonIdx])
	rest := strings.TrimSpace(line[colonIdx+1:])

	rest = strings.TrimPrefix(rest, "{")
	rest = strings.TrimSuffix(rest, "}")
	rest = strings.TrimSpace(rest)

	pd := pipetypes.ParamDef{Name: name}
	for _, part := range strings.Split(rest, ",") {
		kv := strings.SplitN(strings.TrimSpace(part), ":", 2)
		if len(kv) != 2 {
			continue
		}
		key := strings.TrimSpace(kv[0])
		val := strings.Trim(strings.TrimSpace(kv[1]), `"'`)
		switch key {
		case "type":
			pd.DataType = val
		case "default":
			pd.DefaultValue = val
		}
	}
	return pd, nil
}

// extractNodeName derives the node name from a NODE_HEADER token.
//
//	"  @filtered: # comment\n"  →  "filtered"
func extractNodeName(headerText string) string {
	text := strings.TrimSpace(strings.TrimRight(headerText, "\r\n"))
	text = strings.TrimPrefix(text, "@")
	if idx := strings.IndexByte(text, ':'); idx >= 0 {
		text = text[:idx]
	}
	return strings.TrimSpace(text)
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

func dedent(s string) string {
	lines := strings.Split(s, "\n")
	minIndent := -1
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		n := len(line) - len(strings.TrimLeft(line, " \t"))
		if minIndent < 0 || n < minIndent {
			minIndent = n
		}
	}
	if minIndent <= 0 {
		return s
	}
	out := make([]string, len(lines))
	for i, line := range lines {
		if len(line) >= minIndent {
			out[i] = line[minIndent:]
		}
	}
	return strings.Join(out, "\n")
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
