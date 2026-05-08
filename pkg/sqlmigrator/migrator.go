// Package sqlmigrator runs bun migrations against a SQLStore.
package sqlmigrator

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

// Migrator applies and rolls back database migrations.
type Migrator interface {
	Migrate(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type migrator struct {
	m *migrate.Migrator
}

// New creates a Migrator. migrations is the collection built by sqlmigration.New.
func New(db *bun.DB, migrations *migrate.Migrations) Migrator {
	return &migrator{
		m: migrate.NewMigrator(
			db,
			migrations,
			migrate.WithTableName("migration"),
			migrate.WithLocksTableName("migration_lock"),
			migrate.WithMarkAppliedOnSuccess(true),
		),
	}
}

// Migrate initialises the migration tables, acquires the lock, and applies all
// pending migrations. Safe to call on every startup — it is a no-op when the
// database is already up to date.
func (r *migrator) Migrate(ctx context.Context) error {
	if err := r.m.Init(ctx); err != nil {
		return err
	}
	if err := r.lock(ctx); err != nil {
		return err
	}
	defer r.m.Unlock(ctx) //nolint:errcheck

	group, err := r.m.Migrate(ctx)
	if err != nil {
		return err
	}
	if group.IsZero() {
		slog.InfoContext(ctx, "sqlmigrator: database is up to date")
		return nil
	}
	slog.InfoContext(ctx, "sqlmigrator: migrations applied", slog.String("group", group.String()))
	return nil
}

// Rollback rolls back the last migration group.
func (r *migrator) Rollback(ctx context.Context) error {
	if err := r.lock(ctx); err != nil {
		return err
	}
	defer r.m.Unlock(ctx) //nolint:errcheck

	group, err := r.m.Rollback(ctx)
	if err != nil {
		return err
	}
	if group.IsZero() {
		slog.InfoContext(ctx, "sqlmigrator: nothing to roll back")
		return nil
	}
	slog.InfoContext(ctx, "sqlmigrator: rolled back", slog.String("group", group.String()))
	return nil
}

// lock tries to acquire the bun migration lock, retrying every 10s for up to 2m.
func (r *migrator) lock(ctx context.Context) error {
	if err := r.m.Lock(ctx); err == nil {
		return nil
	}

	timer := time.NewTimer(2 * time.Minute)
	defer timer.Stop()
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timer.C:
			return fmt.Errorf("sqlmigrator: timed out waiting for migration lock")
		case <-ticker.C:
			if err := r.m.Lock(ctx); err == nil {
				return nil
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
