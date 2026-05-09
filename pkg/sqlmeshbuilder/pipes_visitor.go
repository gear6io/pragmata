package sqlmeshbuilder

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	pipelang "github.com/gear6io/pragmata/pkg/grammars/pipesgrammar"
	sqlmesh "github.com/gear6io/pragmata/pkg/grammars/sqlmeshgrammar"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

// tok returns the grammar-defined string literal for a SQLMesh token type.
// LiteralNames entries are ANTLR single-quoted (e.g. "'FULL'"); we strip those quotes.
func tok(tokenType int) string {
	return strings.Trim(sqlmesh.SQLMeshParserStaticData.LiteralNames[tokenType], "'")
}

// dialect is the ClickHouse dialect value emitted in every model block.
const dialect = "clickhouse"

// PipesVisitor walks a PipeLang parse tree and emits a SQLMesh model definition.
// It embeds BasePipeLangVisitor so only the directives we care about need overrides.
type PipesVisitor struct {
	pipelang.BasePipeLangVisitor
	name        string
	kind        string // tok(SQLMeshKIND_FULL) or tok(SQLMeshKIND_VIEW)
	modelName   string // defaults to pipe name; overridden by datasource directives
	copyTarget  string // non-empty for COPY pipes — used for INSERT INTO
	cron        string
	description string
	tags        []string
	nodes       []pipetypes.Node
	errs        []error
}

// Visit dispatches to the typed VisitXxx method via Accept.
func (v *PipesVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

// VisitChildren visits all child nodes of a rule node in order.
func (v *PipesVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for i := 0; i < node.GetChildCount(); i++ {
		if pt, ok := node.GetChild(i).(antlr.ParseTree); ok {
			pt.Accept(v)
		}
	}
	return nil
}

// ── Directive visitors ────────────────────────────────────────────────────────

func (v *PipesVisitor) VisitPipeFile(ctx *pipelang.PipeFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *PipesVisitor) VisitStatement(ctx *pipelang.StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *PipesVisitor) VisitDescriptionDir(ctx *pipelang.DescriptionDirContext) interface{} {
	v.description = strings.TrimSpace(ctx.REST_OF_LINE().GetText())
	return nil
}

func (v *PipesVisitor) VisitTagsDir(ctx *pipelang.TagsDirContext) interface{} {
	raw := strings.TrimSpace(ctx.REST_OF_LINE().GetText())
	for _, t := range strings.Split(raw, ",") {
		if tag := strings.TrimSpace(t); tag != "" {
			v.tags = append(v.tags, tag)
		}
	}
	return nil
}

func (v *PipesVisitor) VisitTypeDir(ctx *pipelang.TypeDirContext) interface{} {
	raw := strings.ToLower(strings.TrimSpace(ctx.REST_OF_LINE().GetText()))
	switch raw {
	case "materialized", "copy":
		v.kind = tok(sqlmesh.SQLMeshKIND_FULL)
	default:
		v.kind = tok(sqlmesh.SQLMeshKIND_VIEW)
	}
	return nil
}

func (v *PipesVisitor) VisitDatasourceDir(ctx *pipelang.DatasourceDirContext) interface{} {
	v.modelName = strings.TrimSpace(ctx.REST_OF_LINE().GetText())
	return nil
}

func (v *PipesVisitor) VisitTargetDatasourceDir(ctx *pipelang.TargetDatasourceDirContext) interface{} {
	ds := strings.TrimSpace(ctx.REST_OF_LINE().GetText())
	v.modelName = ds
	v.copyTarget = ds
	return nil
}

func (v *PipesVisitor) VisitCopyScheduleDir(ctx *pipelang.CopyScheduleDirContext) interface{} {
	v.cron = strings.TrimSpace(ctx.REST_OF_LINE().GetText())
	return nil
}

// ── Node / SQL visitors ───────────────────────────────────────────────────────

func (v *PipesVisitor) VisitNodeBlock(ctx *pipelang.NodeBlockContext) interface{} {
	name := strings.TrimSpace(ctx.REST_OF_LINE().GetText())
	if name == "" {
		v.errs = append(v.errs, fmt.Errorf("NODE requires a name"))
		return nil
	}
	result := v.Visit(ctx.SqlBlock())
	node, ok := result.(pipetypes.Node)
	if !ok {
		return nil
	}
	node.Name = name
	v.nodes = append(v.nodes, node)
	return nil
}

func (v *PipesVisitor) VisitSqlBlock(ctx *pipelang.SqlBlockContext) interface{} {
	return v.Visit(ctx.SqlBody())
}

func (v *PipesVisitor) VisitSqlBody(ctx *pipelang.SqlBodyContext) interface{} {
	lines := make([]string, 0, len(ctx.AllSQL_LINE()))
	for _, t := range ctx.AllSQL_LINE() {
		lines = append(lines, strings.TrimRight(t.GetText(), "\r\n"))
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

// ── Output assembly ───────────────────────────────────────────────────────────

// Build validates the collected state and returns the SQLMesh model definition.
func (v *PipesVisitor) Build() (string, error) {
	if len(v.errs) > 0 {
		return "", v.errs[0]
	}
	if len(v.nodes) == 0 {
		return "", fmt.Errorf("pipe %q has no NODE blocks", v.name)
	}
	if v.kind == "" {
		return "", fmt.Errorf("pipe %q has no TYPE declaration", v.name)
	}

	var sb strings.Builder

	// MODEL ( ... )
	sb.WriteString(tok(sqlmesh.SQLMeshMODEL))
	sb.WriteString(" (\n")
	writeProp(&sb, tok(sqlmesh.SQLMeshPROP_NAME), v.modelName)
	writeProp(&sb, tok(sqlmesh.SQLMeshPROP_KIND), v.kind)
	writeProp(&sb, tok(sqlmesh.SQLMeshPROP_DIALECT), quoted(dialect))
	if v.description != "" {
		writeProp(&sb, tok(sqlmesh.SQLMeshPROP_DESCRIPTION), quoted(v.description))
	}
	if len(v.tags) > 0 {
		writeProp(&sb, tok(sqlmesh.SQLMeshPROP_TAGS), tagsArray(v.tags))
	}
	if v.cron != "" {
		writeProp(&sb, tok(sqlmesh.SQLMeshPROP_CRON), quoted(v.cron))
	}
	sb.WriteString(");\n\n")

	// SQL body
	sql := buildSQL(v.nodes, v.copyTarget)
	sb.WriteString(sql)
	if !strings.HasSuffix(sql, "\n") {
		sb.WriteByte('\n')
	}

	return sb.String(), nil
}

// writeProp appends "  key = value,\n" using grammar-derived key names.
func writeProp(sb *strings.Builder, key, value string) {
	fmt.Fprintf(sb, "  %s = %s,\n", key, value)
}

// quoted wraps s in single quotes, escaping any embedded single quotes.
func quoted(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "''") + "'"
}

// tagsArray formats a tag slice as a SQLMesh array literal.
func tagsArray(tags []string) string {
	q := make([]string, len(tags))
	for i, t := range tags {
		q[i] = quoted(t)
	}
	return "[" + strings.Join(q, ", ") + "]"
}

// buildSQL assembles the SQL body from nodes.
// All-but-last nodes become CTEs; the last node's SQL is the final SELECT.
// COPY pipes prepend INSERT INTO <target>.
func buildSQL(nodes []pipetypes.Node, copyTarget string) string {
	var sb strings.Builder

	if copyTarget != "" {
		fmt.Fprintf(&sb, "INSERT INTO %s\n", copyTarget)
	}

	if len(nodes) == 1 {
		sb.WriteString(strings.TrimSpace(nodes[0].SQL))
		return sb.String()
	}

	sb.WriteString("WITH\n")
	for i, node := range nodes[:len(nodes)-1] {
		fmt.Fprintf(&sb, "  %s AS (\n", node.Name)
		for _, line := range strings.Split(strings.TrimSpace(node.SQL), "\n") {
			fmt.Fprintf(&sb, "    %s\n", line)
		}
		sb.WriteString("  )")
		if i < len(nodes)-2 {
			sb.WriteByte(',')
		}
		sb.WriteByte('\n')
	}
	sb.WriteByte('\n')
	sb.WriteString(strings.TrimSpace(nodes[len(nodes)-1].SQL))
	return sb.String()
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

// ── Entry point ───────────────────────────────────────────────────────────────

// FromContent parses raw .pipe file content and returns a SQLMesh .sql string.
func FromContent(pipeName, content string) (string, error) {
	if !strings.HasSuffix(content, "\n") {
		content += "\n"
	}

	lexer := pipelang.NewPipeLangLexer(antlr.NewInputStream(content))
	lexErr := &errListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexErr)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := pipelang.NewPipeLang(stream)
	parseErr := &errListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(parseErr)

	tree := p.PipeFile()

	if lexErr.msg != "" {
		return "", fmt.Errorf("lex error in %q: %s", pipeName, lexErr.msg)
	}
	if parseErr.msg != "" {
		return "", fmt.Errorf("parse error in %q: %s", pipeName, parseErr.msg)
	}

	v := &PipesVisitor{name: pipeName, modelName: pipeName}
	v.Visit(tree)
	return v.Build()
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

func init() {
	// Populate SQLMeshParserStaticData.LiteralNames so tok() works without
	// first constructing a parser instance.
	sqlmesh.SQLMeshInit()
}
