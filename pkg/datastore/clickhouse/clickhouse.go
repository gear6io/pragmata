package clickhouse

import (
	"context"
	"fmt"
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2"
	driver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"

	"github.com/gear6io/pragmata/pkg/types/querybuildertypes"
	"github.com/gear6io/pragmata/pkg/types/sourcetypes"
)

const sourceDatabase = "pragmata_source"

// Store connects to ClickHouse and manages sources in the pragmata_source database.
type Store struct {
	conn driver.Conn
}

// New opens a ClickHouse connection from a URL (e.g. clickhouse://user:pass@host:9000/db).
func New(url string) (*Store, error) {
	opts, err := clickhouse.ParseDSN(url)
	if err != nil {
		return nil, fmt.Errorf("parse clickhouse url: %w", err)
	}
	conn, err := clickhouse.Open(opts)
	if err != nil {
		return nil, fmt.Errorf("open clickhouse: %w", err)
	}
	return &Store{conn: conn}, nil
}

// CreateSource issues a CREATE TABLE in the pragmata_source database.
func (s *Store) CreateSource(ctx context.Context, src *sourcetypes.Source) error {
	if len(src.Fields) == 0 {
		return fmt.Errorf("source must have at least one field")
	}

	engine := src.Engine
	if engine == "" {
		engine = "MergeTree()"
	}

	var colDefs []string
	for _, f := range src.Fields {
		chType, err := f.Type.ClickHouseType()
		if err != nil {
			return fmt.Errorf("field %q: %w", f.Name, err)
		}
		colDefs = append(colDefs, fmt.Sprintf("%s %s", f.Name, chType))
	}

	query := fmt.Sprintf(
		"CREATE TABLE %s.%s (%s) ENGINE = %s",
		sourceDatabase, src.Name, strings.Join(colDefs, ", "), engine,
	)
	if err := s.conn.Exec(ctx, query); err != nil {
		return fmt.Errorf("create source %q: %w", src.Name, err)
	}
	src.Database = sourceDatabase
	return nil
}

// ListSources returns all tables in the pragmata_source database.
func (s *Store) ListSources(ctx context.Context) ([]sourcetypes.Source, error) {
	rows, err := s.conn.Query(ctx,
		"SELECT name, engine FROM system.tables WHERE database = ?",
		sourceDatabase,
	)
	if err != nil {
		return nil, fmt.Errorf("list sources: %w", err)
	}
	defer rows.Close()

	var sources []sourcetypes.Source
	for rows.Next() {
		var src sourcetypes.Source
		if err := rows.Scan(&src.Name, &src.Engine); err != nil {
			return nil, fmt.Errorf("scan source: %w", err)
		}
		src.Database = sourceDatabase
		sources = append(sources, src)
	}
	return sources, rows.Err()
}

// GetSource returns a single table with its field list.
func (s *Store) GetSource(ctx context.Context, name string) (*sourcetypes.Source, error) {
	rows, err := s.conn.Query(ctx,
		"SELECT name, engine FROM system.tables WHERE database = ? AND name = ?",
		sourceDatabase, name,
	)
	if err != nil {
		return nil, fmt.Errorf("get source %q: %w", name, err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}
	var src sourcetypes.Source
	if err := rows.Scan(&src.Name, &src.Engine); err != nil {
		return nil, fmt.Errorf("scan source %q: %w", name, err)
	}
	src.Database = sourceDatabase
	rows.Close()

	fieldRows, err := s.conn.Query(ctx,
		"SELECT name, type FROM system.columns WHERE database = ? AND table = ?",
		sourceDatabase, name,
	)
	if err != nil {
		return nil, fmt.Errorf("get source fields %q: %w", name, err)
	}
	defer fieldRows.Close()

	for fieldRows.Next() {
		var colName, colType string
		if err := fieldRows.Scan(&colName, &colType); err != nil {
			return nil, fmt.Errorf("scan field: %w", err)
		}
		src.Fields = append(src.Fields, querybuildertypes.Field{
			Name: colName,
			Type: querybuildertypes.FieldDataTypeFromClickHouse(colType),
		})
	}
	return &src, fieldRows.Err()
}
