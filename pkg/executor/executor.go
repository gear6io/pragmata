package executor

import (
	"context"

	"github.com/gear6io/pragmata/pkg/types/executortypes"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/gear6io/pragmata/pkg/valuer"
)

// JobStatus is the lifecycle state of an execution job.
type JobStatus struct {
	valuer.String
}

var (
	JobStatusRunning  = JobStatus{valuer.NewString("running")}
	JobStatusFailed   = JobStatus{valuer.NewString("failed")}
	JobStatusComplete = JobStatus{valuer.NewString("complete")}
)

// MaterializedPipeParams carries everything needed to create a materialized pipe.
type MaterializedPipeParams struct {
	Pipe              *pipetypes.ExecutablePipe
	BackfillIntervals []executortypes.TimeInterval
}

// JobID is an opaque string identifying an execution job.
type JobID = string

// Job is the observable state of an execution job.
type Job struct {
	ID                 JobID
	PipeID             string
	Status             JobStatus
	TotalIntervals     int
	CompletedIntervals int
	LastError          string
}

// Executor drives long-running pipe lifecycle operations.
//
// The goroutineexecutor implementation uses goroutines and SQLite checkpoints.
// Swap in a Temporal implementation via dependency injection when needed.
type Executor interface {
	// StartMaterializedPipe writes the SQLMesh model, applies it, and runs the
	// full backfill loop. Returns a JobID for status polling. Non-blocking.
	StartMaterializedPipe(ctx context.Context, params MaterializedPipeParams) (JobID, error)

	// RunPipe fires a single execution of a pipe's SQLMesh model. Blocking.
	// Works for both COPY and MATERIALIZED pipes.
	RunPipe(ctx context.Context, pipeID string) error

	// GetJobStatus returns the current state of a previously started job.
	GetJobStatus(ctx context.Context, jobID JobID) (*Job, error)
}
