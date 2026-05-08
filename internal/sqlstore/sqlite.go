// Package sqlstore provides a SQLite-backed implementation of sqlstore.SQLStore.
package sqlstore

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "modernc.org/sqlite" // pure-Go SQLite driver

	"github.com/gear6io/pragmata/pkg/sqlstore"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

// SQLiteStore implements sqlstore.SQLStore using SQLite via database/sql.
type SQLiteStore struct {
	db *sql.DB
}

// New opens (or creates) a SQLite database at path and runs schema migrations.
func New(path string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open sqlite %q: %w", path, err)
	}
	db.SetMaxOpenConns(1) // SQLite is single-writer
	if err := migrate(db); err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}
	return &SQLiteStore{db: db}, nil
}

func migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS pipes (
			name       TEXT PRIMARY KEY,
			type       TEXT NOT NULL,
			data       TEXT NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS backfill_jobs (
			id                   TEXT PRIMARY KEY,
			pipe_id              TEXT NOT NULL,
			total_intervals      INTEGER NOT NULL DEFAULT 0,
			completed_intervals  INTEGER NOT NULL DEFAULT 0,
			status               TEXT NOT NULL DEFAULT 'running',
			last_error           TEXT NOT NULL DEFAULT '',
			updated_at           DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS tokens (
			id         TEXT PRIMARY KEY,
			name       TEXT NOT NULL,
			value      TEXT UNIQUE NOT NULL,
			scope      TEXT NOT NULL DEFAULT 'PIPES:READ',
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`)
	return err
}

// --- Pipe methods ---

func (s *SQLiteStore) CreatePipe(ctx context.Context, pipe *pipetypes.Pipe) error {
	data, err := json.Marshal(pipe)
	if err != nil {
		return err
	}
	_, err = s.db.ExecContext(ctx,
		`INSERT INTO pipes (name, type, data) VALUES (?, ?, ?)`,
		pipe.Name, string(pipe.Type), string(data),
	)
	return err
}

func (s *SQLiteStore) GetPipe(ctx context.Context, name string) (*pipetypes.Pipe, error) {
	row := s.db.QueryRowContext(ctx, `SELECT data FROM pipes WHERE name = ?`, name)
	var data string
	if err := row.Scan(&data); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("pipe %q not found", name)
		}
		return nil, err
	}
	var pipe pipetypes.Pipe
	return &pipe, json.Unmarshal([]byte(data), &pipe)
}

func (s *SQLiteStore) ListPipes(ctx context.Context) ([]*pipetypes.Pipe, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT data FROM pipes ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var pipes []*pipetypes.Pipe
	for rows.Next() {
		var data string
		if err := rows.Scan(&data); err != nil {
			return nil, err
		}
		var p pipetypes.Pipe
		if err := json.Unmarshal([]byte(data), &p); err != nil {
			return nil, err
		}
		pipes = append(pipes, &p)
	}
	return pipes, rows.Err()
}

func (s *SQLiteStore) UpdatePipe(ctx context.Context, pipe *pipetypes.Pipe) error {
	data, err := json.Marshal(pipe)
	if err != nil {
		return err
	}
	res, err := s.db.ExecContext(ctx,
		`UPDATE pipes SET type = ?, data = ?, updated_at = CURRENT_TIMESTAMP WHERE name = ?`,
		string(pipe.Type), string(data), pipe.Name,
	)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return fmt.Errorf("pipe %q not found", pipe.Name)
	}
	return nil
}

func (s *SQLiteStore) DeletePipe(ctx context.Context, name string) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM pipes WHERE name = ?`, name)
	return err
}

// --- BackfillJob methods ---

func (s *SQLiteStore) CreateBackfillJob(ctx context.Context, job *sqlstore.BackfillJob) error {
	_, err := s.db.ExecContext(ctx,
		`INSERT INTO backfill_jobs (id, pipe_id, total_intervals, completed_intervals, status, last_error, updated_at)
		 VALUES (?, ?, ?, ?, ?, ?, ?)`,
		job.ID, job.PipeID, job.TotalIntervals, job.CompletedIntervals,
		string(job.Status), job.LastError, job.UpdatedAt,
	)
	return err
}

func (s *SQLiteStore) UpdateBackfillJob(ctx context.Context, job *sqlstore.BackfillJob) error {
	_, err := s.db.ExecContext(ctx,
		`UPDATE backfill_jobs SET completed_intervals = ?, status = ?, last_error = ?, updated_at = ?
		 WHERE id = ?`,
		job.CompletedIntervals, string(job.Status), job.LastError, job.UpdatedAt, job.ID,
	)
	return err
}

func (s *SQLiteStore) GetBackfillJob(ctx context.Context, id string) (*sqlstore.BackfillJob, error) {
	row := s.db.QueryRowContext(ctx,
		`SELECT id, pipe_id, total_intervals, completed_intervals, status, last_error, updated_at
		 FROM backfill_jobs WHERE id = ?`, id)
	return scanBackfillJob(row)
}

func (s *SQLiteStore) ListRunningBackfillJobs(ctx context.Context) ([]*sqlstore.BackfillJob, error) {
	rows, err := s.db.QueryContext(ctx,
		`SELECT id, pipe_id, total_intervals, completed_intervals, status, last_error, updated_at
		 FROM backfill_jobs WHERE status = 'running'`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var jobs []*sqlstore.BackfillJob
	for rows.Next() {
		job, err := scanBackfillJob(rows)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, rows.Err()
}

type scanner interface {
	Scan(dest ...any) error
}

func scanBackfillJob(s scanner) (*sqlstore.BackfillJob, error) {
	var j sqlstore.BackfillJob
	var status string
	var updatedAt string
	err := s.Scan(&j.ID, &j.PipeID, &j.TotalIntervals, &j.CompletedIntervals,
		&status, &j.LastError, &updatedAt)
	if err != nil {
		return nil, err
	}
	j.Status = sqlstore.BackfillStatus(status)
	j.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)
	return &j, nil
}

// --- Token methods ---

func (s *SQLiteStore) CreateToken(ctx context.Context, token *sqlstore.Token) error {
	_, err := s.db.ExecContext(ctx,
		`INSERT INTO tokens (id, name, value, scope, created_at) VALUES (?, ?, ?, ?, ?)`,
		token.ID, token.Name, token.Value, string(token.Scope), token.CreatedAt,
	)
	return err
}

func (s *SQLiteStore) GetTokenByValue(ctx context.Context, value string) (*sqlstore.Token, error) {
	row := s.db.QueryRowContext(ctx,
		`SELECT id, name, value, scope, created_at FROM tokens WHERE value = ?`, value)
	return scanToken(row)
}

func (s *SQLiteStore) ListTokens(ctx context.Context) ([]*sqlstore.Token, error) {
	rows, err := s.db.QueryContext(ctx,
		`SELECT id, name, value, scope, created_at FROM tokens ORDER BY created_at`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tokens []*sqlstore.Token
	for rows.Next() {
		t, err := scanToken(rows)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, t)
	}
	return tokens, rows.Err()
}

func (s *SQLiteStore) DeleteToken(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM tokens WHERE id = ?`, id)
	return err
}

func scanToken(s scanner) (*sqlstore.Token, error) {
	var t sqlstore.Token
	var scope, createdAt string
	err := s.Scan(&t.ID, &t.Name, &t.Value, &scope, &createdAt)
	if err != nil {
		return nil, err
	}
	t.Scope = sqlstore.TokenScope(scope)
	t.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	return &t, nil
}
