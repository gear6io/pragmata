// Package sqlmesh provides Go wrappers around the SQLMesh CLI.
//
// Phase B stubs — all methods return nil without executing anything.
// Replace each stub with real os/exec calls as Phase B is implemented.
package sqlmesh

import (
	"context"

	"github.com/gear6io/pragmata/pkg/types/orchestratortypes"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

// Runner wraps sqlmesh CLI invocations.
type Runner struct {
	ProjectDir string
	BinaryPath string // defaults to "sqlmesh"
}

// New creates a Runner. BinaryPath defaults to "sqlmesh" if empty.
func New(projectDir, binaryPath string) *Runner {
	if binaryPath == "" {
		binaryPath = "sqlmesh"
	}
	return &Runner{ProjectDir: projectDir, BinaryPath: binaryPath}
}

// SyncModel writes (or updates) the SQLMesh .sql model file for pipe.
// TODO Phase B: render model template → write to ProjectDir/models/<name>.sql
func (r *Runner) SyncModel(_ context.Context, _ *pipetypes.Pipe) error {
	return nil
}

// RemoveModel deletes the SQLMesh .sql model file for the given pipe name.
// TODO Phase B: delete ProjectDir/models/<name>.sql
func (r *Runner) RemoveModel(_ context.Context, _ string) error {
	return nil
}

// PlanApply runs `sqlmesh plan --auto-apply` in ProjectDir.
// TODO Phase B: exec.CommandContext → capture stdout/stderr → parse exit code
func (r *Runner) PlanApply(_ context.Context) error {
	return nil
}

// Run executes `sqlmesh run [--model name] [--start S --end E]`.
// interval may be nil for a full run (used by COPY pipes).
// TODO Phase B: exec.CommandContext
func (r *Runner) Run(_ context.Context, _ string, _ *orchestratortypes.TimeInterval) error {
	return nil
}

// BackfillIntervals generates the list of daily intervals from the pipe's
// configured start date to today. Used by the orchestrator on resume.
// TODO Phase B: read start date from config.yaml or pipe metadata.
func (r *Runner) BackfillIntervals(_ *pipetypes.Pipe) []orchestratortypes.TimeInterval {
	return nil
}
