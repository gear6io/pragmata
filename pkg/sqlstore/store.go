package sqlstore

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/gear6io/pragmata/pkg/types/sqlstoretypes"
)

// SQLStore persists pipe definitions, backfill job state, and auth tokens.
type SQLStore interface {
	// BunDB returns the underlying bun.DB, used by the migrator.
	BunDB() *bun.DB

	// Pipes
	CreatePipe(ctx context.Context, pipe *pipetypes.StorablePipe) error
	GetPipe(ctx context.Context, name string) (*pipetypes.Pipe, error)
	ListPipes(ctx context.Context) ([]*pipetypes.Pipe, error)
	UpdatePipe(ctx context.Context, pipe *pipetypes.StorablePipe) error
	DeletePipe(ctx context.Context, name string) error

	// BackfillJobs track long-running SQLMesh materialization progress.
	CreateBackfillJob(ctx context.Context, job *sqlstoretypes.BackfillJob) error
	UpdateBackfillJob(ctx context.Context, job *sqlstoretypes.BackfillJob) error
	GetBackfillJob(ctx context.Context, id string) (*sqlstoretypes.BackfillJob, error)
	ListRunningBackfillJobs(ctx context.Context) ([]*sqlstoretypes.BackfillJob, error)

	// Tokens are bearer tokens for API authentication.
	CreateToken(ctx context.Context, token *sqlstoretypes.Token) error
	GetTokenByValue(ctx context.Context, value string) (*sqlstoretypes.Token, error)
	ListTokens(ctx context.Context) ([]*sqlstoretypes.Token, error)
	DeleteToken(ctx context.Context, id string) error
}
