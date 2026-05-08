package sqlstoretypes

import (
	"time"

	"github.com/gear6io/pragmata/pkg/valuer"
	"github.com/uptrace/bun"
)

// BackfillStatus is the lifecycle state of a backfill job.
type BackfillStatus struct {
	valuer.String
}

var (
	BackfillStatusRunning  = BackfillStatus{valuer.NewString("running")}
	BackfillStatusFailed   = BackfillStatus{valuer.NewString("failed")}
	BackfillStatusComplete = BackfillStatus{valuer.NewString("complete")}
)

// BackfillJob tracks incremental SQLMesh materialization for a MATERIALIZED pipe.
type BackfillJob struct {
	bun.BaseModel      `bun:"table:backfill_jobs"`
	ID                 string         `bun:"id,pk"               json:"id"`
	PipeID             string         `bun:"pipe_id"             json:"pipeId"`
	TotalIntervals     int            `bun:"total_intervals"     json:"totalIntervals"`
	CompletedIntervals int            `bun:"completed_intervals" json:"completedIntervals"`
	Status             BackfillStatus `bun:"status"              json:"status"`
	LastError          string         `bun:"last_error"          json:"lastError"`
	UpdatedAt          time.Time      `bun:"updated_at"          json:"updatedAt"`
}

// TokenScope defines what operations a token may perform.
type TokenScope struct {
	valuer.String
}

var (
	TokenScopeRead  = TokenScope{valuer.NewString("PIPES:READ")}
	TokenScopeWrite = TokenScope{valuer.NewString("PIPES:WRITE")}
)

// Token is a bearer token for API auth.
type Token struct {
	bun.BaseModel `bun:"table:tokens"`
	ID            string     `bun:"id,pk"      json:"id"`
	Name          string     `bun:"name"       json:"name"`
	Value         string     `bun:"value"      json:"value"`
	Scope         TokenScope `bun:"scope"      json:"scope"`
	CreatedAt     time.Time  `bun:"created_at" json:"createdAt"`
}
