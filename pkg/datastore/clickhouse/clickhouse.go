package clickhouse

import (
	"context"
	"fmt"
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2"
	driver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/huandu/go-sqlbuilder"

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

func (s *Store) createTableStmt(src *sourcetypes.Source) (*sqlbuilder.CreateTableBuilder, error) {
	if len(src.Fields) == 0 {
		return nil, fmt.Errorf("source must have at least one field")
	}

	engine := src.Engine
	if engine == sourcetypes.EngineUndefined {
		engine = sourcetypes.EngineMergeTree
	}

	ctb := sqlbuilder.NewCreateTableBuilder()
	ctb.IfNotExists()
	ctb.CreateTable(sourceDatabase + "." + src.Name)
	orderBy := make([]string, 0, len(src.Fields))
	for _, f := range src.Fields {
		chType, err := f.Type.ClickHouseType()
		if err != nil {
			return nil, fmt.Errorf("field %q: %w", f.Name, err)
		}
		ctb.Define(f.Name, chType)
		orderBy = append(orderBy, f.Name)
	}
	ctb.Define("__attrs__", "JSON")
	ctb.Option("ENGINE = "+string(engine.String()), "ORDER BY ("+strings.Join(orderBy, ",")+")")
	return ctb, nil
}

// CreateSource issues a CREATE TABLE in the pragmata_source database.
func (s *Store) CreateSource(ctx context.Context, src *sourcetypes.Source) error {
	ctb, err := s.createTableStmt(src)
	if err != nil {
		return err
	}

	query, args := ctb.BuildWithFlavor(sqlbuilder.ClickHouse)
	if err := s.conn.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("create source %q: %w", src.Name, err)
	}
	src.Database = sourceDatabase
	return nil
}

// ListSources returns all sources in the pragmata_source database.
func (s *Store) ListSources(ctx context.Context) ([]sourcetypes.Source, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("name", "engine")
	sb.From("system.tables")
	sb.Where(sb.Equal("database", sourceDatabase))
	query, args := sb.Build()

	rows, err := s.conn.Query(ctx, query, args...)
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
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("name", "engine")
	sb.From("system.tables")
	sb.Where(sb.Equal("database", sourceDatabase), sb.Equal("name", name))
	query, args := sb.Build()

	rows, err := s.conn.Query(ctx, query, args...)
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

	fsb := sqlbuilder.NewSelectBuilder()
	fsb.Select("name", "type")
	fsb.From("system.columns")
	fsb.Where(fsb.Equal("database", sourceDatabase), fsb.Equal("table", name))
	fieldQuery, fieldArgs := fsb.Build()

	fieldRows, err := s.conn.Query(ctx, fieldQuery, fieldArgs...)
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
