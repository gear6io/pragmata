package orchestration

import (
	"context"
	"time"

	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

// JobStatus is the lifecycle state of an orchestration job.
type JobStatus string

const (
	JobStatusRunning  JobStatus = "running"
	JobStatusFailed   JobStatus = "failed"
	JobStatusComplete JobStatus = "complete"
)

// TimeInterval is a half-open time range [Start, End) used for backfill batching.
type TimeInterval struct {
	Start time.Time
	End   time.Time
}

// MaterializedPipeParams carries everything needed to create a materialized pipe.
type MaterializedPipeParams struct {
	Pipe              *pipetypes.Pipe
	BackfillIntervals []TimeInterval
}

// JobID is an opaque string identifying an orchestration job.
type JobID = string

// Job is the observable state of an orchestration job.
type Job struct {
	ID                 JobID
	PipeID             string
	Status             JobStatus
	TotalIntervals     int
	CompletedIntervals int
	LastError          string
}

// Orchestrator drives long-running pipe lifecycle operations.
//
// The interface is designed so the Phase 1 goroutine implementation and a future
// Temporal implementation are interchangeable via dependency injection.
type Orchestrator interface {
	// StartMaterializedPipe writes the SQLMesh model, applies it, and runs the
	// full backfill loop. Returns a JobID for status polling. Non-blocking.
	StartMaterializedPipe(ctx context.Context, params MaterializedPipeParams) (JobID, error)

	// RunCopyPipe fires a single execution of a COPY pipe's SQL. Blocking.
	RunCopyPipe(ctx context.Context, pipeID string) error

	// GetJobStatus returns the current state of a previously started job.
	GetJobStatus(ctx context.Context, jobID JobID) (*Job, error)
}
