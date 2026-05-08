package sqlstore

import (
	"context"
	"time"

	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

// SQLStore persists pipe definitions, backfill job state, and auth tokens.
type SQLStore interface {
	// Pipes
	CreatePipe(ctx context.Context, pipe *pipetypes.Pipe) error
	GetPipe(ctx context.Context, name string) (*pipetypes.Pipe, error)
	ListPipes(ctx context.Context) ([]*pipetypes.Pipe, error)
	UpdatePipe(ctx context.Context, pipe *pipetypes.Pipe) error
	DeletePipe(ctx context.Context, name string) error

	// BackfillJobs track long-running SQLMesh materialization progress.
	CreateBackfillJob(ctx context.Context, job *BackfillJob) error
	UpdateBackfillJob(ctx context.Context, job *BackfillJob) error
	GetBackfillJob(ctx context.Context, id string) (*BackfillJob, error)
	ListRunningBackfillJobs(ctx context.Context) ([]*BackfillJob, error)

	// Tokens are bearer tokens for API authentication.
	CreateToken(ctx context.Context, token *Token) error
	GetTokenByValue(ctx context.Context, value string) (*Token, error)
	ListTokens(ctx context.Context) ([]*Token, error)
	DeleteToken(ctx context.Context, id string) error
}

// BackfillStatus is the lifecycle state of a backfill job.
type BackfillStatus string

const (
	BackfillStatusRunning  BackfillStatus = "running"
	BackfillStatusFailed   BackfillStatus = "failed"
	BackfillStatusComplete BackfillStatus = "complete"
)

// BackfillJob tracks incremental SQLMesh materialization for a MATERIALIZED pipe.
// Persisted so the goroutine orchestrator can resume after a process restart.
type BackfillJob struct {
	ID                 string
	PipeID             string
	TotalIntervals     int
	CompletedIntervals int
	Status             BackfillStatus
	LastError          string
	UpdatedAt          time.Time
}

// TokenScope defines what operations a token may perform.
type TokenScope string

const (
	TokenScopeRead  TokenScope = "PIPES:READ"
	TokenScopeWrite TokenScope = "PIPES:WRITE"
)

// Token is a bearer token for API auth.
type Token struct {
	ID        string
	Name      string
	Value     string
	Scope     TokenScope
	CreatedAt time.Time
}
