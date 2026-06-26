// Package pipevisitor converts raw .pipe content into a Pipe using the ANTLR4-generated grammar.
// Run scripts/generate-pipe-grammar.sh to regenerate grammar/ after editing the .g4 files.
package pipevisitor

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	grammar "github.com/gear6io/pragmata/pkg/grammars/pipesgrammar"
	"github.com/gear6io/pragmata/pkg/prqlvisitor"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/gear6io/pragmata/pkg/types/sourcetypes"
	"github.com/gear6io/pragmata/pkg/valuer"
)

type PipeVisitorOpts struct {
	FetchSources    func(srcs ...string) ([]sourcetypes.Source, error)
	SourceValidator prqlvisitor.SourceValidator
	FieldValidator  prqlvisitor.Validator
}

func (opts *PipeVisitorOpts) validate() error {
	if opts.SourceValidator != nil && opts.FetchSources == nil {
		return fmt.Errorf("FetchSources can not be nil, with SourceValidator")
	}

	return nil
}

// pipeVisitor implements grammar.PipeLangVisitor, accumulating a Pipe and any errors.
type pipeVisitor struct {
	grammar.BasePipeLangVisitor
	pipe         *pipetypes.Pipe
	validSources []sourcetypes.Source
	errs         []error
	opts         PipeVisitorOpts
}

// Visit converts raw .pipe file content into a Pipe. name is the initial pipe
// name (typically derived from the filename); a name: directive in the file
// overrides it.
func Visit(name, content string, opts PipeVisitorOpts) (*pipetypes.Pipe, error) {
	if err := opts.validate(); err != nil {
		return nil, err
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

	v := &pipeVisitor{
		pipe: &pipetypes.Pipe{Name: name, Content: content},
		opts: opts,
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
	v.pipe.Type = mapPipeType(ctx.VALUE().GetText())
	return nil
}

func (v *pipeVisitor) VisitName(ctx *grammar.NameContext) interface{} {
	v.pipe.Name = ctx.VALUE().GetText()
	return nil
}

func (v *pipeVisitor) VisitDescription(ctx *grammar.DescriptionContext) interface{} {
	v.pipe.Description = ctx.VALUE().GetText()
	return nil
}

func (v *pipeVisitor) VisitDescriptionML(ctx *grammar.DescriptionMLContext) interface{} {
	lines := make([]string, 0, len(ctx.AllSECTION_LINE()))
	for _, tok := range ctx.AllSECTION_LINE() {
		if s := strings.TrimSpace(tok.GetText()); s != "" {
			lines = append(lines, s)
		}
	}
	v.pipe.Description = strings.Join(lines, " ")
	return nil
}

func (v *pipeVisitor) VisitTags(ctx *grammar.TagsContext) interface{} {
	raw := strings.Trim(ctx.VALUE().GetText(), "[] \t")
	for _, tag := range strings.Split(raw, ",") {
		if t := strings.TrimSpace(tag); t != "" {
			v.pipe.Tags = append(v.pipe.Tags, t)
		}
	}
	return nil
}

func (v *pipeVisitor) VisitOwner(ctx *grammar.OwnerContext) interface{} {
	v.pipe.Owner = ctx.VALUE().GetText()
	return nil
}

func (v *pipeVisitor) VisitDestination(ctx *grammar.DestinationContext) interface{} {
	v.pipe.Destination = ctx.VALUE().GetText()
	return nil
}

func (v *pipeVisitor) VisitSchedule(ctx *grammar.ScheduleContext) interface{} {
	v.pipe.Schedule = ctx.VALUE().GetText()
	return nil
}

func (v *pipeVisitor) VisitSourcesClause(ctx *grammar.SourcesClauseContext) interface{} {
	for _, e := range ctx.AllSource() {
		e.(antlr.ParseTree).Accept(v)
	}

	// validations
	// Check for no two same alias
	{
		sourcesSet := map[string]struct{}{}
		for _, src := range v.pipe.Sources {
			alias := src.Table
			if src.Alias != "" {
				alias = src.Alias
			}
			_, found := sourcesSet[alias]
			if found {
				return fmt.Errorf("alias collision in sources")
			}
			sourcesSet[alias] = struct{}{}
		}
	}
	// Check for valid sources
	if v.opts.SourceValidator != nil {
		sources := []string{}
		for _, src := range v.pipe.Sources {
			sources = append(sources, src.Table)
		}
		var err error
		v.validSources, err = v.opts.FetchSources(sources...)
		if err != nil {
			return err
		}

		for _, src := range v.pipe.Sources {
			if err := v.opts.SourceValidator(src.Table, v.validSources); err != nil {
				v.errs = append(v.errs, err)
			}
		}
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
	var nodes []pipetypes.Node
	for _, tok := range ctx.AllSECTION_LINE() {
		line := strings.TrimSpace(tok.GetText())
		if line == "" {
			continue
		}
		if rest, ok := strings.CutPrefix(line, "@"); ok {
			header, _, _ := strings.Cut(rest, ":")
			name := strings.TrimSpace(header)
			if name == "" {
				v.errs = append(v.errs, fmt.Errorf("pipeline node has empty name"))
				return nil
			}
			nodes = append(nodes, pipetypes.Node{Name: name})
		} else if len(nodes) > 0 {
			nodes[len(nodes)-1].SQL = strings.TrimSpace(nodes[len(nodes)-1].SQL + " " + line)
		}
	}

	for _, node := range nodes {
		if v.opts.SourceValidator != nil {
			if _, err := prqlvisitor.Visit(node.SQL, prqlvisitor.PRQLVisitorOpts{
				FromValidator: v.prqlFROMValidator(),
			}); err != nil {
				v.errs = append(v.errs, fmt.Errorf("node %q: %w", node.Name, err))
			}
		}
		v.pipe.Nodes = append(v.pipe.Nodes, node)
	}
	return nil
}

func (v *pipeVisitor) prqlFROMValidator() prqlvisitor.FromValidator {
	return func(table string, isNode bool) error {
		if isNode {
			has := slices.ContainsFunc(v.pipe.Nodes, func(node pipetypes.Node) bool {
				return node.Name == table
			})
			if has {
				return nil
			}
		}

		has := slices.ContainsFunc(v.pipe.Sources, func(src pipetypes.Source) bool {
			return src.String() == table
		})
		if has {
			return nil
		}

		return fmt.Errorf("unknown reference [%s]", table)
	}
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
