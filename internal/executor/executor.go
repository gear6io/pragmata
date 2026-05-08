// Package executor builds and runs ClickHouse queries from pipe node chains.
package executor

import (
	"context"
	"fmt"
	"strings"

	"github.com/gear6io/pragmata/internal/template"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

// Executor runs a SQL query and returns structured results.
type Executor interface {
	Query(ctx context.Context, sql string) (*pipetypes.ExecuteResult, error)
	Close() error
}

// BuildCTE assembles a single SQL statement from pipe nodes by wrapping
// intermediate nodes as CTEs. The last node becomes the outermost SELECT.
//
//	WITH node1 AS (sql1),
//	     node2 AS (sql2)
//	<last_node_sql>
//
// Template placeholders in nodes with IsTemplated=true are rendered with params.
func BuildCTE(nodes []pipetypes.Node, params map[string]string) (string, error) {
	if len(nodes) == 0 {
		return "", fmt.Errorf("BuildCTE: no nodes")
	}

	// Render templates for each node.
	type rendered struct{ name, sql string }
	rNodes := make([]rendered, len(nodes))
	for i, n := range nodes {
		sql := n.SQL
		if n.IsTemplated {
			var err error
			sql, err = template.Render(sql, params)
			if err != nil {
				return "", fmt.Errorf("node %q: %w", n.Name, err)
			}
		}
		rNodes[i] = rendered{n.Name, sql}
	}

	// Single node: return its SQL directly.
	if len(rNodes) == 1 {
		return rNodes[0].sql, nil
	}

	// Multi-node: intermediate nodes → CTEs, last node → final query.
	var sb strings.Builder
	sb.WriteString("WITH ")
	for i, n := range rNodes[:len(rNodes)-1] {
		if i > 0 {
			sb.WriteString(",\n")
		}
		sb.WriteString(n.name)
		sb.WriteString(" AS (\n")
		sb.WriteString(indent(n.sql, "  "))
		sb.WriteString("\n)")
	}
	sb.WriteString("\n")
	sb.WriteString(rNodes[len(rNodes)-1].sql)

	return sb.String(), nil
}

func indent(s, prefix string) string {
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		if l != "" {
			lines[i] = prefix + l
		}
	}
	return strings.Join(lines, "\n")
}
