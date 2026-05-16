package sqlmeshbuilder

import (
	"fmt"
	"strings"

	sqlmesh "github.com/gear6io/pragmata/pkg/grammars/sqlmeshgrammar"
	"github.com/gear6io/pragmata/pkg/parser/pipeparser"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/huandu/go-sqlbuilder"
)

// tok returns the grammar-defined string literal for a SQLMesh token type.
// LiteralNames entries are ANTLR single-quoted (e.g. "'FULL'"); we strip those quotes.
func tok(tokenType int) string {
	return strings.Trim(sqlmesh.SQLMeshParserStaticData.LiteralNames[tokenType], "'")
}

// dialect is the ClickHouse dialect value emitted in every model block.
const dialect = "clickhouse"

// FromContent parses raw .pipe file content and returns a SQLMesh .sql string.
func FromPipe(pipe *pipetypes.Pipe) (string, error) {
	pipe, err := pipeparser.Parse(pipe.Name, pipe.Content)
	if err != nil {
		return "", err
	}

	kind := kindForType(pipe.Type)
	modelName := coalesce(pipe.Destination, pipe.Name)

	var sb strings.Builder

	// MODEL ( ... )
	sb.WriteString(tok(sqlmesh.SQLMeshMODEL))
	sb.WriteString(" (\n")
	writeProp(&sb, tok(sqlmesh.SQLMeshPROP_NAME), modelName)
	writeProp(&sb, tok(sqlmesh.SQLMeshPROP_KIND), kind)
	writeProp(&sb, tok(sqlmesh.SQLMeshPROP_DIALECT), quoted(dialect))
	if pipe.Description != "" {
		writeProp(&sb, tok(sqlmesh.SQLMeshPROP_DESCRIPTION), quoted(pipe.Description))
	}
	if len(pipe.Tags) > 0 {
		writeProp(&sb, tok(sqlmesh.SQLMeshPROP_TAGS), tagsArray(pipe.Tags))
	}
	if pipe.Schedule != "" {
		writeProp(&sb, tok(sqlmesh.SQLMeshPROP_CRON), quoted(pipe.Schedule))
	}
	sb.WriteString(");\n\n")

	// SQL body — copyTarget is empty for the new format (destination is the model name).
	sql, err := buildSQL(pipe.Nodes, "")
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
		pipetypes.PipeTypeTable, pipetypes.PipeTypeIncremental, pipetypes.PipeTypeSnapshot:
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
		return "", fmt.Errorf("no nodes provided")
	}

	outBuilder := sqlbuilder.NewInsertBuilder()
	if copyTarget != "" {
		outBuilder.InsertInto(copyTarget)
	}

	lastNode := nodes[len(nodes)-1]
	lastSB, err := QueryBuilder(lastNode.SQL)
	if err != nil {
		return "", fmt.Errorf("node %q: %w", lastNode.Name, err)
	}

	for _, node := range nodes[:len(nodes)-1] {
		sb, err := QueryBuilder(node.SQL)
		if err != nil {
			return "", fmt.Errorf("node %q: %w", node.Name, err)
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
