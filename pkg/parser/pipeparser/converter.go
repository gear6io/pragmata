package pipeparser

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gear6io/pragmata/pkg/parser/pipeparser/grammar"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/gear6io/pragmata/pkg/valuer"
)

// converter walks the PipeLang parse tree and builds a *pipetypes.Pipe.
// It embeds BasePipeLangVisitor so only the rules we care about need overrides.
type converter struct {
	grammar.BasePipeLangVisitor
	pipe *pipetypes.Pipe
	errs []error
}

// Visit dispatches to the typed VisitXxx method via Accept.
func (c *converter) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(c)
}

// VisitChildren visits all child nodes of a rule node in order.
func (c *converter) VisitChildren(node antlr.RuleNode) interface{} {
	for i := 0; i < node.GetChildCount(); i++ {
		if pt, ok := node.GetChild(i).(antlr.ParseTree); ok {
			pt.Accept(c)
		}
	}
	return nil
}

// result validates the constructed Pipe and returns it, or the first error.
func (c *converter) result() (*pipetypes.Pipe, error) {
	if len(c.errs) > 0 {
		return nil, c.errs[0]
	}
	return c.pipe, c.validate()
}

func (c *converter) validate() error {
	p := c.pipe
	if len(p.Nodes) == 0 {
		return fmt.Errorf("pipe %q has no NODE blocks", p.Name)
	}
	if p.Type == pipetypes.PipeTypeUndefined {
		return fmt.Errorf("pipe %q has no TYPE declaration", p.Name)
	}
	if p.Type == pipetypes.PipeTypeMaterialized && p.Datasource == "" {
		return fmt.Errorf("pipe %q: TYPE MATERIALIZED requires DATASOURCE", p.Name)
	}
	if p.Type == pipetypes.PipeTypeCopy && p.TargetDatasource == "" {
		return fmt.Errorf("pipe %q: TYPE COPY requires TARGET_DATASOURCE", p.Name)
	}
	return nil
}

// ── Visitor overrides ─────────────────────────────────────────────────────────

func (c *converter) VisitPipeFile(ctx *grammar.PipeFileContext) interface{} {
	return c.VisitChildren(ctx)
}

func (c *converter) VisitStatement(ctx *grammar.StatementContext) interface{} {
	return c.VisitChildren(ctx)
}

func (c *converter) VisitDescriptionDir(ctx *grammar.DescriptionDirContext) interface{} {
	c.pipe.Description = strings.TrimSpace(ctx.REST_OF_LINE().GetText())
	return nil
}

func (c *converter) VisitTagsDir(ctx *grammar.TagsDirContext) interface{} {
	raw := strings.TrimSpace(ctx.REST_OF_LINE().GetText())
	for _, t := range strings.Split(raw, ",") {
		if tag := strings.TrimSpace(t); tag != "" {
			c.pipe.Tags = append(c.pipe.Tags, tag)
		}
	}
	return nil
}

func (c *converter) VisitTypeDir(ctx *grammar.TypeDirContext) interface{} {
	raw := strings.TrimSpace(ctx.REST_OF_LINE().GetText())
	c.pipe.Type = pipetypes.PipeType{String: valuer.NewString(raw)}
	return nil
}

func (c *converter) VisitDatasourceDir(ctx *grammar.DatasourceDirContext) interface{} {
	c.pipe.Datasource = strings.TrimSpace(ctx.REST_OF_LINE().GetText())
	return nil
}

func (c *converter) VisitTargetDatasourceDir(ctx *grammar.TargetDatasourceDirContext) interface{} {
	c.pipe.TargetDatasource = strings.TrimSpace(ctx.REST_OF_LINE().GetText())
	return nil
}

func (c *converter) VisitCopyScheduleDir(ctx *grammar.CopyScheduleDirContext) interface{} {
	c.pipe.CopySchedule = strings.TrimSpace(ctx.REST_OF_LINE().GetText())
	return nil
}

func (c *converter) VisitNodeBlock(ctx *grammar.NodeBlockContext) interface{} {
	name := strings.TrimSpace(ctx.REST_OF_LINE().GetText())
	if name == "" {
		c.errs = append(c.errs, fmt.Errorf("NODE requires a name"))
		return nil
	}
	result := c.Visit(ctx.SqlBlock())
	node, ok := result.(pipetypes.Node)
	if !ok {
		return nil
	}
	node.Name = name
	c.pipe.Nodes = append(c.pipe.Nodes, node)
	return nil
}

func (c *converter) VisitSqlBlock(ctx *grammar.SqlBlockContext) interface{} {
	return c.Visit(ctx.SqlBody())
}

func (c *converter) VisitSqlBody(ctx *grammar.SqlBodyContext) interface{} {
	// Collect raw line texts. Each SQL_LINE token includes its trailing newline.
	lines := make([]string, 0, len(ctx.AllSQL_LINE()))
	for _, tok := range ctx.AllSQL_LINE() {
		lines = append(lines, strings.TrimRight(tok.GetText(), "\r\n"))
	}
	raw := strings.Join(lines, "\n")

	isTemplated := false
	trimmedRaw := strings.TrimLeft(raw, "\n\r ")
	if strings.HasPrefix(trimmedRaw, "%") {
		afterPercent := strings.TrimPrefix(trimmedRaw, "%")
		switch {
		case strings.HasPrefix(afterPercent, "\n"):
			raw = afterPercent[1:]
		case afterPercent == "":
			raw = ""
		default:
			raw = afterPercent
		}
		isTemplated = true
	}

	return pipetypes.Node{
		SQL:         strings.TrimSpace(dedent(raw)),
		IsTemplated: isTemplated,
	}
}

// dedent removes the common leading whitespace from all non-empty lines.
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
