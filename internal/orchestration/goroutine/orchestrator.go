// Package goroutine provides a goroutine-based implementation of orchestration.Orchestrator.
//
// Long-running operations (SQLMesh apply + backfill) run in background goroutines.
// Progress is checkpointed in SQLite after every interval so a process restart
// can resume from where it left off via ResumeInterrupted.
package goroutine

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/gear6io/pragmata/internal/sqlmesh"
	"github.com/gear6io/pragmata/pkg/orchestration"
	"github.com/gear6io/pragmata/pkg/sqlstore"
)

// Orchestrator implements orchestration.Orchestrator using goroutines and SQLite checkpoints.
type Orchestrator struct {
	store  sqlstore.SQLStore
	runner *sqlmesh.Runner
}

// New creates an Orchestrator. Call ResumeInterrupted after startup to continue
// any backfill jobs that were in-progress when the process last stopped.
func New(store sqlstore.SQLStore, runner *sqlmesh.Runner) *Orchestrator {
	return &Orchestrator{store: store, runner: runner}
}

// ResumeInterrupted restarts goroutines for any backfill jobs left in "running"
// state by a previous process. Call once on startup, before serving requests.
func (o *Orchestrator) ResumeInterrupted(ctx context.Context) error {
	jobs, err := o.store.ListRunningBackfillJobs(ctx)
	if err != nil {
		return fmt.Errorf("list running jobs: %w", err)
	}
	for _, job := range jobs {
		pipe, err := o.store.GetPipe(ctx, job.PipeID)
		if err != nil {
			// Pipe was deleted while job was running; mark it failed.
			job.Status = sqlstore.BackfillStatusFailed
			job.LastError = "pipe not found on resume"
			job.UpdatedAt = time.Now()
			_ = o.store.UpdateBackfillJob(ctx, job)
			continue
		}
		// Re-generate intervals; skip already completed ones.
		intervals := o.runner.BackfillIntervals(pipe)
		remaining := intervals[job.CompletedIntervals:]
		go o.runBackfill(job, pipe.Name, remaining)
	}
	return nil
}

// StartMaterializedPipe creates a backfill job and begins the SQLMesh lifecycle
// in a background goroutine. Returns the job ID for status polling.
func (o *Orchestrator) StartMaterializedPipe(ctx context.Context, params orchestration.MaterializedPipeParams) (orchestration.JobID, error) {
	jobID := uuid.New().String()
	job := &sqlstore.BackfillJob{
		ID:             jobID,
		PipeID:         params.Pipe.Name,
		TotalIntervals: len(params.BackfillIntervals),
		Status:         sqlstore.BackfillStatusRunning,
		UpdatedAt:      time.Now(),
	}
	if err := o.store.CreateBackfillJob(ctx, job); err != nil {
		return "", fmt.Errorf("create backfill job: %w", err)
	}

	// Convert orchestration.TimeInterval to sqlmesh.TimeInterval (same fields).
	intervals := make([]sqlmesh.TimeInterval, len(params.BackfillIntervals))
	for i, iv := range params.BackfillIntervals {
		intervals[i] = sqlmesh.TimeInterval{Start: iv.Start, End: iv.End}
	}

	go o.runMaterializedPipeline(job, params.Pipe.Name, intervals)
	return jobID, nil
}

// RunCopyPipe fires a single sqlmesh run for a COPY pipe. Blocking.
func (o *Orchestrator) RunCopyPipe(ctx context.Context, pipeID string) error {
	return o.runner.Run(ctx, pipeID, nil)
}

// GetJobStatus returns the current state of a job.
func (o *Orchestrator) GetJobStatus(ctx context.Context, jobID orchestration.JobID) (*orchestration.Job, error) {
	j, err := o.store.GetBackfillJob(ctx, jobID)
	if err != nil {
		return nil, err
	}
	return &orchestration.Job{
		ID:                 j.ID,
		PipeID:             j.PipeID,
		Status:             orchestration.JobStatus(j.Status),
		TotalIntervals:     j.TotalIntervals,
		CompletedIntervals: j.CompletedIntervals,
		LastError:          j.LastError,
	}, nil
}

// runMaterializedPipeline is the full lifecycle: model sync → plan/apply → backfill.
func (o *Orchestrator) runMaterializedPipeline(job *sqlstore.BackfillJob, pipeID string, intervals []sqlmesh.TimeInterval) {
	ctx := context.Background()

	pipe, err := o.store.GetPipe(ctx, pipeID)
	if err != nil {
		o.failJob(ctx, job, fmt.Errorf("load pipe: %w", err))
		return
	}

	if err := o.runner.SyncModel(ctx, pipe); err != nil {
		o.failJob(ctx, job, fmt.Errorf("sync model: %w", err))
		return
	}
	if err := o.runner.PlanApply(ctx); err != nil {
		o.failJob(ctx, job, fmt.Errorf("sqlmesh plan apply: %w", err))
		return
	}

	o.runBackfill(job, pipeID, intervals)
}

// runBackfill iterates over intervals, calling sqlmesh run for each, with checkpointing.
func (o *Orchestrator) runBackfill(job *sqlstore.BackfillJob, pipeID string, intervals []sqlmesh.TimeInterval) {
	ctx := context.Background()
	for _, iv := range intervals {
		if err := o.runner.Run(ctx, pipeID, &iv); err != nil {
			o.failJob(ctx, job, err)
			return
		}
		job.CompletedIntervals++
		job.UpdatedAt = time.Now()
		_ = o.store.UpdateBackfillJob(ctx, job) // best-effort checkpoint
	}
	job.Status = sqlstore.BackfillStatusComplete
	job.UpdatedAt = time.Now()
	_ = o.store.UpdateBackfillJob(ctx, job)
}

func (o *Orchestrator) failJob(ctx context.Context, job *sqlstore.BackfillJob, err error) {
	job.Status = sqlstore.BackfillStatusFailed
	job.LastError = err.Error()
	job.UpdatedAt = time.Now()
	_ = o.store.UpdateBackfillJob(ctx, job)
}
