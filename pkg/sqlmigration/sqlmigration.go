package sqlmigration

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

// SQLMigration is the interface for a single versioned schema change.
// Each implementation must live in its own .go file so bun can derive
// the migration name from the call site of Register.
type SQLMigration interface {
	Register(*migrate.Migrations) error
	Up(context.Context, *bun.DB) error
	Down(context.Context, *bun.DB) error
}

// New registers all provided migrations in order and returns the bun Migrations
// collection ready to be handed to sqlmigrator.
func New(migrations []SQLMigration) (*migrate.Migrations, error) {
	ms := migrate.NewMigrations()
	for _, m := range migrations {
		if err := m.Register(ms); err != nil {
			return nil, err
		}
	}
	return ms, nil
}
