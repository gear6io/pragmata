// Package querier executes ENDPOINT pipes directly against ClickHouse.
// It owns: param template substitution → PRQL compilation → CTE assembly → execution → row scanning.
package querier

import (
	"context"
	"regexp"
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2"
	driver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/huandu/go-sqlbuilder"

	"github.com/gear6io/pragmata/pkg/errors"
	"github.com/gear6io/pragmata/pkg/prqlvisitor"
	"github.com/gear6io/pragmata/pkg/template"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

const sourceDatabase = "pragmata_source"

// Querier executes ENDPOINT pipes against ClickHouse.
type Querier struct {
	conn driver.Conn
}

// New opens a dedicated ClickHouse connection for query execution.
func New(url string) (*Querier, error) {
	opts, err := clickhouse.ParseDSN(url)
	if err != nil {
		return nil, errors.WrapInvalidInputf(err, errors.CodeInvalidInput, "parse clickhouse url")
	}
	conn, err := clickhouse.Open(opts)
	if err != nil {
		return nil, errors.WrapInternalf(err, errors.CodeInternal, "open clickhouse")
	}
	return &Querier{conn: conn}, nil
}

// Execute resolves params, compiles PRQL nodes, assembles a CTE chain with source aliases,
// runs the query, and returns the rows.
func (q *Querier) Execute(ctx context.Context, pipe *pipetypes.ExecutablePipe, urlParams map[string]string) (*pipetypes.ExecuteResult, error) {
	if len(pipe.Nodes) == 0 {
		return nil, errors.NewInvalidInputf(errors.CodeInvalidInput, "pipe %q has no nodes", pipe.Name)
	}

	// Render template tokens in each node's PRQL then compile to SelectBuilder.
	sbs := make([]*sqlbuilder.SelectBuilder, len(pipe.Nodes))
	for i, node := range pipe.Nodes {
		rendered, err := renderNode(node.SQL, pipe.Params, urlParams)
		if err != nil {
			return nil, errors.WrapInternalf(err, errors.CodeInternal, "node %q", node.Name)
		}
		sb, err := prqlvisitor.Visit(rendered, prqlvisitor.PRQLVisitorOpts{})
		if err != nil {
			return nil, errors.WrapInternalf(err, errors.CodeInternal, "node %q: compile", node.Name)
		}
		sbs[i] = sb
	}

	last := len(pipe.Nodes) - 1
	root := sbs[last]

	// Source CTEs must come first so compiled node SQL can reference them.
	for _, src := range pipe.Sources {
		srcSB := sqlbuilder.NewSelectBuilder().Select("*").From(sourceDatabase + "." + src.Table)
		root.With(sqlbuilder.With(sqlbuilder.CTEQuery(src.Alias).As(srcSB)))
	}

	// Preceding nodes become CTEs in order.
	for i, node := range pipe.Nodes[:last] {
		root.With(sqlbuilder.With(sqlbuilder.CTEQuery(node.Name).As(sbs[i])))
	}

	sql, _ := root.BuildWithFlavor(sqlbuilder.ClickHouse)

	rows, err := q.conn.Query(ctx, sql)
	if err != nil {
		return nil, errors.WrapInternalf(err, errors.CodeInternal, "execute")
	}
	defer rows.Close()

	data, err := scanRows(rows)
	if err != nil {
		return nil, err
	}
	return &pipetypes.ExecuteResult{Data: data}, nil
}

// paramTokenRe matches {{ params.name }} with optional whitespace.
var paramTokenRe = regexp.MustCompile(`\{\{\s*params\.(\w+)\s*\}\}`)

// renderNode expands {{ params.X }} tokens using ParamDefs (type-aware SQL formatting),
// then delegates remaining {{ Type(name, default) }} tokens to template.Render.
func renderNode(prql string, defs pipetypes.ParamDefs, urlParams map[string]string) (string, error) {
	// Normalize {{ params.X }} → {{ Type(X, default) }} using declared param metadata.
	defMap := make(map[string]pipetypes.ParamDef, len(defs))
	for _, d := range defs {
		defMap[d.Name] = d
	}
	normalized := paramTokenRe.ReplaceAllStringFunc(prql, func(match string) string {
		sub := paramTokenRe.FindStringSubmatch(match)
		name := sub[1]
		def, ok := defMap[name]
		if !ok {
			return match // unknown param — leave for template.Render to handle or fail
		}
		typeName, err := def.DataType.ClickHouseType()
		if err != nil {
			return match
		}
		return "{{ " + typeName + "(" + name + ", " + quoteDefault(def.DefaultValue) + ") }}"
	})
	return template.Render(normalized, urlParams)
}

// quoteDefault wraps a default value in single quotes for the template engine String/DateTime types.
func quoteDefault(v string) string {
	if v == "" {
		return "''"
	}
	return "'" + strings.ReplaceAll(v, "'", "''") + "'"
}

// scanRows reads all rows into a slice of column-name → value maps.
func scanRows(rows driver.Rows) ([]map[string]any, error) {
	cols := rows.Columns()
	var result []map[string]any
	for rows.Next() {
		vals := make([]any, len(cols))
		ptrs := make([]any, len(cols))
		for i := range vals {
			ptrs[i] = &vals[i]
		}
		if err := rows.Scan(ptrs...); err != nil {
			return nil, errors.WrapInternalf(err, errors.CodeInternal, "scan row")
		}
		row := make(map[string]any, len(cols))
		for i, col := range cols {
			row[col] = vals[i]
		}
		result = append(result, row)
	}
	if result == nil {
		result = []map[string]any{}
	}
	return result, rows.Err()
}
