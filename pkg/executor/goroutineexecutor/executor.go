// Package goroutineexecutor provides a goroutine-based implementation of executor.Executor.
//
// Long-running operations (SQLMesh apply + backfill) run in background goroutines.
// Progress is checkpointed in SQLite after every interval so a process restart
// can resume from where it left off via ResumeInterrupted.
package goroutineexecutor

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/gear6io/pragmata/pkg/executor"
	"github.com/gear6io/pragmata/pkg/pipevisitor"
	"github.com/gear6io/pragmata/pkg/sqlmesh"
	"github.com/gear6io/pragmata/pkg/sqlstore"
	"github.com/gear6io/pragmata/pkg/types/executortypes"
	"github.com/gear6io/pragmata/pkg/types/sqlstoretypes"
)

// Executor implements executor.Executor using goroutines and SQLite checkpoints.
type Executor struct {
	store  sqlstore.SQLStore
	runner *sqlmesh.Runner
}

// New creates an Executor. Call ResumeInterrupted after startup to continue
// any backfill jobs that were in-progress when the process last stopped.
func New(store sqlstore.SQLStore, runner *sqlmesh.Runner) *Executor {
	return &Executor{store: store, runner: runner}
}

// ResumeInterrupted restarts goroutines for any backfill jobs left in "running"
// state by a previous process. Call once on startup, before serving requests.
func (e *Executor) ResumeInterrupted(ctx context.Context) error {
	jobs, err := e.store.ListRunningBackfillJobs(ctx)
	if err != nil {
		return fmt.Errorf("list running jobs: %w", err)
	}
	for _, job := range jobs {
		pipe, err := e.store.GetPipe(ctx, job.PipeID)
		if err != nil {
			// Pipe was deleted while job was running; mark it failed.
			job.Status = sqlstoretypes.BackfillStatusFailed
			job.LastError = "pipe not found on resume"
			job.UpdatedAt = time.Now()
			_ = e.store.UpdateBackfillJob(ctx, job)
			continue
		}
		// Re-generate intervals; skip already completed ones.
		intervals := e.runner.BackfillIntervals(&pipe.Pipe)
		remaining := intervals[job.CompletedIntervals:]
		go e.runBackfill(job, pipe.Name, remaining)
	}
	return nil
}

// StartMaterializedPipe creates a backfill job and begins the SQLMesh lifecycle
// in a background goroutine. Returns the job ID for status polling.
func (e *Executor) StartMaterializedPipe(ctx context.Context, params executor.MaterializedPipeParams) (executor.JobID, error) {
	jobID := uuid.New().String()
	job := &sqlstoretypes.BackfillJob{
		ID:             jobID,
		PipeID:         params.Pipe.Name,
		TotalIntervals: len(params.BackfillIntervals),
		Status:         sqlstoretypes.BackfillStatusRunning,
		UpdatedAt:      time.Now(),
	}
	if err := e.store.CreateBackfillJob(ctx, job); err != nil {
		return "", fmt.Errorf("create backfill job: %w", err)
	}

	intervals := make([]executortypes.TimeInterval, len(params.BackfillIntervals))
	for i, iv := range params.BackfillIntervals {
		intervals[i] = executortypes.TimeInterval{Start: iv.Start, End: iv.End}
	}

	go e.runMaterializedPipeline(job, params.Pipe.Name, intervals)
	return jobID, nil
}

// RunPipe fires a single sqlmesh run for a pipe. Blocking.
// Resolves the SQLMesh model name from the stored pipe content.
func (e *Executor) RunPipe(ctx context.Context, pipeID string) error {
	pipe, err := e.store.GetPipe(ctx, pipeID)
	if err != nil {
		return fmt.Errorf("load pipe %q: %w", pipeID, err)
	}
	exec, err := pipevisitor.Visit(pipe.Content, pipevisitor.PipeVisitorOpts{})
	if err != nil {
		return fmt.Errorf("parse pipe %q: %w", pipeID, err)
	}
	modelName := coalesce(exec.Destination, exec.Name)
	return e.runner.Run(ctx, modelName, nil)
}

// GetJobStatus returns the current state of a job.
func (e *Executor) GetJobStatus(ctx context.Context, jobID executor.JobID) (*executor.Job, error) {
	j, err := e.store.GetBackfillJob(ctx, jobID)
	if err != nil {
		return nil, err
	}
	return &executor.Job{
		ID:                 j.ID,
		PipeID:             j.PipeID,
		Status:             executor.JobStatus(j.Status),
		TotalIntervals:     j.TotalIntervals,
		CompletedIntervals: j.CompletedIntervals,
		LastError:          j.LastError,
	}, nil
}

// runMaterializedPipeline is the full lifecycle: model sync → plan/apply → backfill.
func (e *Executor) runMaterializedPipeline(job *sqlstoretypes.BackfillJob, pipeID string, intervals []executortypes.TimeInterval) {
	ctx := context.Background()

	pipe, err := e.store.GetPipe(ctx, pipeID)
	if err != nil {
		e.failJob(ctx, job, fmt.Errorf("load pipe: %w", err))
		return
	}

	if err := e.runner.SyncModel(ctx, &pipe.Pipe); err != nil {
		e.failJob(ctx, job, fmt.Errorf("sync model: %w", err))
		return
	}
	if err := e.runner.PlanApply(ctx); err != nil {
		e.failJob(ctx, job, fmt.Errorf("sqlmesh plan apply: %w", err))
		return
	}

	e.runBackfill(job, pipeID, intervals)
}

// runBackfill iterates over intervals, calling sqlmesh run for each, with checkpointing.
func (e *Executor) runBackfill(job *sqlstoretypes.BackfillJob, pipeID string, intervals []executortypes.TimeInterval) {
	ctx := context.Background()
	for _, iv := range intervals {
		if err := e.runner.Run(ctx, pipeID, &iv); err != nil {
			e.failJob(ctx, job, err)
			return
		}
		job.CompletedIntervals++
		job.UpdatedAt = time.Now()
		_ = e.store.UpdateBackfillJob(ctx, job) // best-effort checkpoint
	}
	job.Status = sqlstoretypes.BackfillStatusComplete
	job.UpdatedAt = time.Now()
	_ = e.store.UpdateBackfillJob(ctx, job)
}

func (e *Executor) failJob(ctx context.Context, job *sqlstoretypes.BackfillJob, err error) {
	job.Status = sqlstoretypes.BackfillStatusFailed
	job.LastError = err.Error()
	job.UpdatedAt = time.Now()
	_ = e.store.UpdateBackfillJob(ctx, job)
}

func coalesce(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}
