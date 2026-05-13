package sqlmeshbuilder

import (
	"fmt"
	"strings"

	sqlmesh "github.com/gear6io/pragmata/pkg/grammars/sqlmeshgrammar"
	"github.com/gear6io/pragmata/pkg/parser/pipeparser"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

// tok returns the grammar-defined string literal for a SQLMesh token type.
// LiteralNames entries are ANTLR single-quoted (e.g. "'FULL'"); we strip those quotes.
func tok(tokenType int) string {
	return strings.Trim(sqlmesh.SQLMeshParserStaticData.LiteralNames[tokenType], "'")
}

// dialect is the ClickHouse dialect value emitted in every model block.
const dialect = "clickhouse"

// FromContent parses raw .pipe file content and returns a SQLMesh .sql string.
func FromContent(pipeName, content string) (string, error) {
	pipe, err := pipeparser.Parse(pipeName, content)
	if err != nil {
		return "", err
	}

	kind := kindForType(pipe.Type)
	modelName := coalesce(pipe.Destination, pipeName)

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
	sql := buildSQL(pipe.Nodes, "")
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

// buildSQL assembles the SQL body from nodes.
// All-but-last nodes become CTEs; the last node's SQL is the final SELECT.
// copyTarget, when non-empty, prepends INSERT INTO <target>.
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

// errListener is kept for parity; SQLMesh grammar errors surface via pipeparser.
type errListener struct {
	msg string
}

func init() {
	// Populate SQLMeshParserStaticData.LiteralNames so tok() works without
	// first constructing a parser instance.
	sqlmesh.SQLMeshInit()
}
