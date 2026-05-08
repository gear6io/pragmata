// Package sqlitesqlstore provides a SQLite-backed implementation of sqlstore.SQLStore.
package sqlitesqlstore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	_ "modernc.org/sqlite" // pure-Go SQLite driver

	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/gear6io/pragmata/pkg/types/sqlstoretypes"
)

// Store implements sqlstore.SQLStore using bun over SQLite.
type Store struct {
	bundb *bun.DB
}

// New opens (or creates) a SQLite database at path.
// Schema migrations are handled separately via pkg/sqlmigrator — call Migrate before serving.
func New(path string) (*Store, error) {
	sqldb, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open sqlite %q: %w", path, err)
	}
	sqldb.SetMaxOpenConns(1) // SQLite is single-writer
	return &Store{bundb: bun.NewDB(sqldb, sqlitedialect.New())}, nil
}

// BunDB returns the underlying bun.DB, used by the migrator.
func (s *Store) BunDB() *bun.DB {
	return s.bundb
}

// --- Pipe methods ---

func (s *Store) CreatePipe(ctx context.Context, pipe *pipetypes.Pipe) error {
	_, err := s.bundb.NewInsert().Model(pipe).Exec(ctx)
	return err
}

func (s *Store) GetPipe(ctx context.Context, name string) (*pipetypes.Pipe, error) {
	pipe := new(pipetypes.Pipe)
	if err := s.bundb.NewSelect().Model(pipe).Where("name = ?", name).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("pipe %q not found", name)
		}
		return nil, err
	}
	return pipe, nil
}

func (s *Store) ListPipes(ctx context.Context) ([]*pipetypes.Pipe, error) {
	var pipes []*pipetypes.Pipe
	if err := s.bundb.NewSelect().Model(&pipes).OrderExpr("name ASC").Scan(ctx); err != nil {
		return nil, err
	}
	return pipes, nil
}

func (s *Store) UpdatePipe(ctx context.Context, pipe *pipetypes.Pipe) error {
	pipe.UpdatedAt = time.Now()
	res, err := s.bundb.NewUpdate().Model(pipe).
		Column("type", "description", "tags", "nodes", "datasource", "target_datasource", "copy_schedule", "updated_at").
		WherePK().
		Exec(ctx)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("pipe %q not found", pipe.Name)
	}
	return nil
}

func (s *Store) DeletePipe(ctx context.Context, name string) error {
	_, err := s.bundb.NewDelete().Model((*pipetypes.Pipe)(nil)).Where("name = ?", name).Exec(ctx)
	return err
}

// --- BackfillJob methods ---

func (s *Store) CreateBackfillJob(ctx context.Context, job *sqlstoretypes.BackfillJob) error {
	_, err := s.bundb.NewInsert().Model(job).Exec(ctx)
	return err
}

func (s *Store) UpdateBackfillJob(ctx context.Context, job *sqlstoretypes.BackfillJob) error {
	_, err := s.bundb.NewUpdate().Model(job).
		Column("completed_intervals", "status", "last_error", "updated_at").
		WherePK().
		Exec(ctx)
	return err
}

func (s *Store) GetBackfillJob(ctx context.Context, id string) (*sqlstoretypes.BackfillJob, error) {
	job := new(sqlstoretypes.BackfillJob)
	if err := s.bundb.NewSelect().Model(job).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return job, nil
}

func (s *Store) ListRunningBackfillJobs(ctx context.Context) ([]*sqlstoretypes.BackfillJob, error) {
	var jobs []*sqlstoretypes.BackfillJob
	if err := s.bundb.NewSelect().Model(&jobs).
		Where("status = ?", sqlstoretypes.BackfillStatusRunning).
		Scan(ctx); err != nil {
		return nil, err
	}
	return jobs, nil
}

// --- Token methods ---

func (s *Store) CreateToken(ctx context.Context, token *sqlstoretypes.Token) error {
	_, err := s.bundb.NewInsert().Model(token).Exec(ctx)
	return err
}

func (s *Store) GetTokenByValue(ctx context.Context, value string) (*sqlstoretypes.Token, error) {
	token := new(sqlstoretypes.Token)
	if err := s.bundb.NewSelect().Model(token).Where("value = ?", value).Scan(ctx); err != nil {
		return nil, err
	}
	return token, nil
}

func (s *Store) ListTokens(ctx context.Context) ([]*sqlstoretypes.Token, error) {
	var tokens []*sqlstoretypes.Token
	if err := s.bundb.NewSelect().Model(&tokens).OrderExpr("created_at ASC").Scan(ctx); err != nil {
		return nil, err
	}
	return tokens, nil
}

func (s *Store) DeleteToken(ctx context.Context, id string) error {
	_, err := s.bundb.NewDelete().Model((*sqlstoretypes.Token)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
