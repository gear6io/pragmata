package clickhouseexecutor

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/ClickHouse/clickhouse-go/v2" // register clickhouse driver

	"github.com/gear6io/pragmata/pkg/config"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

// Executor runs SQL queries against a ClickHouse instance via database/sql.
type Executor struct {
	db *sql.DB
}

// New opens a connection pool to ClickHouse using the given config.
func New(cfg config.ClickHouseConfig) (*Executor, error) {
	dsn := fmt.Sprintf("clickhouse://%s:%s@%s:%d/%s?dial_timeout=5s&read_timeout=60s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		return nil, fmt.Errorf("open clickhouse: %w", err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)
	return &Executor{db: db}, nil
}

// Query executes sql and returns all rows with metadata and timing stats.
func (e *Executor) Query(ctx context.Context, query string) (*pipetypes.ExecuteResult, error) {
	start := time.Now()

	rows, err := e.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("clickhouse query: %w", err)
	}
	defer rows.Close()

	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, fmt.Errorf("column types: %w", err)
	}
	cols := make([]string, len(colTypes))
	columns := make([]pipetypes.Column, len(colTypes))
	for i, ct := range colTypes {
		cols[i] = ct.Name()
		columns[i] = pipetypes.Column{Name: ct.Name(), Type: ct.DatabaseTypeName()}
	}

	var data []map[string]any
	for rows.Next() {
		vals := make([]any, len(cols))
		ptrs := make([]any, len(cols))
		for i := range vals {
			ptrs[i] = &vals[i]
		}
		if err := rows.Scan(ptrs...); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		row := make(map[string]any, len(cols))
		for i, col := range cols {
			row[col] = vals[i]
		}
		data = append(data, row)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows: %w", err)
	}

	return &pipetypes.ExecuteResult{
		Data: data,
		Meta: pipetypes.ResultMeta{Columns: columns},
		Stats: pipetypes.ResultStats{
			Elapsed:  time.Since(start).Seconds(),
			RowsRead: uint64(len(data)),
		},
	}, nil
}

func (e *Executor) Close() error {
	return e.db.Close()
}
