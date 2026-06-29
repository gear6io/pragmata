package sqlmeshbuilder

import (
	"fmt"
	"strings"

	sqlmesh "github.com/gear6io/pragmata/pkg/grammars/sqlmeshgrammar"
	"github.com/gear6io/pragmata/pkg/errors"
	"github.com/gear6io/pragmata/pkg/pipevisitor"
	"github.com/gear6io/pragmata/pkg/prqlvisitor"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/huandu/go-sqlbuilder"
)

var CodeInvalidPipeContent = errors.MustNewCode("invalid_pipe_content")

// tok returns the grammar-defined string literal for a SQLMesh token type.
// LiteralNames entries are ANTLR single-quoted (e.g. "'FULL'"); we strip those quotes.
func tok(tokenType int) string {
	return strings.Trim(sqlmesh.SQLMeshParserStaticData.LiteralNames[tokenType], "'")
}

// dialect is the ClickHouse dialect value emitted in every model block.
const dialect = "clickhouse"

// FromPipe parses the pipe's Content and returns a SQLMesh .sql string.
func FromPipe(pipe *pipetypes.Pipe) (string, error) {
	exec, err := pipevisitor.Visit(pipe.Content, pipevisitor.PipeVisitorOpts{})
	if err != nil {
		return "", err
	}

	kind := kindForType(exec.Type)
	modelName := coalesce(exec.Destination, exec.Name)

	var sb strings.Builder

	// MODEL ( ... )
	sb.WriteString(tok(sqlmesh.SQLMeshMODEL))
	sb.WriteString(" (\n")
	writeProp(&sb, tok(sqlmesh.SQLMeshPROP_NAME), modelName)
	writeProp(&sb, tok(sqlmesh.SQLMeshPROP_KIND), kind)
	writeProp(&sb, tok(sqlmesh.SQLMeshPROP_DIALECT), quoted(dialect))
	if exec.Description != "" {
		writeProp(&sb, tok(sqlmesh.SQLMeshPROP_DESCRIPTION), quoted(exec.Description))
	}
	if len(exec.Tags) > 0 {
		writeProp(&sb, tok(sqlmesh.SQLMeshPROP_TAGS), tagsArray(exec.Tags))
	}
	if exec.Schedule != "" {
		writeProp(&sb, tok(sqlmesh.SQLMeshPROP_CRON), quoted(exec.Schedule))
	}
	sb.WriteString(");\n\n")

	// SQL body — copyTarget is empty for the new format (destination is the model name).
	sql, err := buildSQL(exec.Nodes, "")
	if err != nil {
		return "", err
	}
	sb.WriteString(sql)
	if !strings.HasSuffix(sql, "\n") {
		sb.WriteByte('\n')
	}

	return sb.String(), nil
}

// kindForType maps a PipeType to a SQLMesh KIND value.
func kindForType(t pipetypes.PipeType) string {
	switch t {
	case pipetypes.PipeTypeMaterialized, pipetypes.PipeTypeCopy,
		pipetypes.PipeTypeTable, pipetypes.PipeTypeSnapshot:
		return tok(sqlmesh.SQLMeshKIND_FULL)
	case pipetypes.PipeTypeIncremental:
		// ponytail: FULL for now; extend to INCREMENTAL_BY_TIME_RANGE when
		// Pipe gains a time_column field in the grammar
		return tok(sqlmesh.SQLMeshKIND_FULL)
	default:
		return tok(sqlmesh.SQLMeshKIND_VIEW)
	}
}

// coalesce returns the first non-empty string.
func coalesce(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
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

// buildSQL compiles each node's PQL source to SQL, then assembles the body.
// All-but-last nodes become CTEs; the last node's SQL is the final SELECT.
// copyTarget, when non-empty, prepends INSERT INTO <target>.
func buildSQL(nodes []pipetypes.Node, copyTarget string) (string, error) {
	if len(nodes) == 0 {
		return "", errors.NewInvalidInputf(CodeInvalidPipeContent, "no nodes provided")
	}

	outBuilder := sqlbuilder.NewInsertBuilder()
	if copyTarget != "" {
		outBuilder.InsertInto(copyTarget)
	}

	lastNode := nodes[len(nodes)-1]
	lastSB, err := prqlvisitor.Visit(lastNode.SQL, prqlvisitor.PRQLVisitorOpts{})
	if err != nil {
		return "", errors.WrapInvalidInputf(err, CodeInvalidPipeContent, "node %q", lastNode.Name)
	}

	for _, node := range nodes[:len(nodes)-1] {
		sb, err := prqlvisitor.Visit(node.SQL, prqlvisitor.PRQLVisitorOpts{})
		if err != nil {
			return "", errors.WrapInvalidInputf(err, CodeInvalidPipeContent, "node %q", node.Name)
		}
		lastSB.With(sqlbuilder.With(sqlbuilder.CTEQuery(node.Name).As(sb)))
	}

	compiled, _ := lastSB.BuildWithFlavor(sqlbuilder.ClickHouse)
	if copyTarget != "" {
		outBuilder.SQL(compiled)

		finalSQL, _ := outBuilder.BuildWithFlavor(sqlbuilder.ClickHouse)
		return finalSQL, nil
	}

	return compiled, nil
}

func init() {
	// Populate SQLMeshParserStaticData.LiteralNames so tok() works without
	// first constructing a parser instance.
	sqlmesh.SQLMeshInit()
}
