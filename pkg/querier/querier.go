// Package querier executes ENDPOINT pipes directly against ClickHouse.
// It owns: param template substitution → PRQL compilation → CTE assembly → execution → row scanning.
package querier

import (
	"context"
	"reflect"
	"regexp"
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2"
	driver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"

	"github.com/gear6io/pragmata/pkg/errors"
	"github.com/gear6io/pragmata/pkg/prqlvisitor"
	"github.com/gear6io/pragmata/pkg/template"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

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

// BuildSQL resolves params, compiles PRQL nodes, and assembles a CTE chain with source aliases,
// returning the final ClickHouse SQL string. It does not touch the database.
func BuildSQL(pipe *pipetypes.ExecutablePipe, urlParams map[string]string) (string, error) {
	if len(pipe.Nodes) == 0 {
		return "", errors.NewInvalidInputf(errors.CodeInvalidInput, "pipe %q has no nodes", pipe.Name)
	}
	rendered := make(pipetypes.Nodes, len(pipe.Nodes))
	for i, node := range pipe.Nodes {
		sql, err := renderNode(node.SQL, pipe.Params, urlParams)
		if err != nil {
			return "", errors.WithAdditionalf(err, "node %q", node.Name)
		}
		rendered[i] = pipetypes.Node{Name: node.Name, SQL: sql}
	}
	return prqlvisitor.BuildSQL(rendered, pipe.Sources)
}

// Execute resolves params, compiles PRQL nodes, assembles a CTE chain with source aliases,
// runs the query, and returns the rows.
func (q *Querier) Execute(ctx context.Context, pipe *pipetypes.ExecutablePipe, urlParams map[string]string) (*pipetypes.ExecuteResult, error) {
	sql, err := BuildSQL(pipe, urlParams)
	if err != nil {
		return nil, err
	}

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
// It uses ColumnTypes().ScanType() to allocate typed destinations so that
// ClickHouse native-protocol types (Date, DateTime64, etc.) scan correctly.
func scanRows(rows driver.Rows) ([]map[string]any, error) {
	cols := rows.Columns()
	colTypes := rows.ColumnTypes()
	var result []map[string]any
	for rows.Next() {
		ptrs := make([]any, len(cols))
		for i, ct := range colTypes {
			ptrs[i] = reflect.New(ct.ScanType()).Interface()
		}
		if err := rows.Scan(ptrs...); err != nil {
			return nil, errors.WrapInternalf(err, errors.CodeInternal, "scan row")
		}
		row := make(map[string]any, len(cols))
		for i, col := range cols {
			row[col] = reflect.ValueOf(ptrs[i]).Elem().Interface()
		}
		result = append(result, row)
	}
	if result == nil {
		result = []map[string]any{}
	}
	return result, rows.Err()
}
