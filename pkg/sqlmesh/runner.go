// Package sqlmesh provides Go wrappers around the SQLMesh CLI.
package sqlmesh

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gear6io/pragmata/pkg/sqlmeshbuilder"
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
func (r *Runner) SyncModel(_ context.Context, pipe *pipetypes.Pipe) error {
	sql, err := sqlmeshbuilder.FromPipe(pipe)
	if err != nil {
		return fmt.Errorf("generate model for %q: %w", pipe.Name, err)
	}
	modelsDir := filepath.Join(r.ProjectDir, "models")
	if err := os.MkdirAll(modelsDir, 0o755); err != nil {
		return fmt.Errorf("create models dir: %w", err)
	}
	dest := filepath.Join(modelsDir, pipe.Name+".sql")
	return os.WriteFile(dest, []byte(sql), 0o644)
}

// RemoveModel deletes the SQLMesh .sql model file for the given pipe name.
func (r *Runner) RemoveModel(_ context.Context, name string) error {
	dest := filepath.Join(r.ProjectDir, "models", name+".sql")
	if err := os.Remove(dest); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("remove model %q: %w", dest, err)
	}
	return nil
}

// PlanApply runs `sqlmesh plan --auto-apply` in ProjectDir.
func (r *Runner) PlanApply(ctx context.Context) error {
	return r.run(ctx, "plan", "--auto-apply")
}

// Run executes `sqlmesh run --model name [--start S --end E]`.
// interval may be nil for a full run (used by COPY pipes).
func (r *Runner) Run(ctx context.Context, pipeID string, interval *orchestratortypes.TimeInterval) error {
	args := []string{"run", "--model", pipeID}
	if interval != nil {
		args = append(args,
			"--start", interval.Start.Format("2006-01-02"),
			"--end", interval.End.Format("2006-01-02"),
		)
	}
	return r.run(ctx, args...)
}

// BackfillIntervals generates the list of daily intervals from the pipe's
// configured start date to today.
// TODO: decide where the start date comes from (per-pipe field or global config)
// and implement date generation. Returns nil until then — the orchestrator
// handles an empty slice gracefully (no backfill runs are scheduled).
func (r *Runner) BackfillIntervals(_ *pipetypes.Pipe) []orchestratortypes.TimeInterval {
	return nil
}

// run executes the sqlmesh binary with the given args inside ProjectDir.
// A non-zero exit code becomes an error that includes captured stdout+stderr.
func (r *Runner) run(ctx context.Context, args ...string) error {
	cmd := exec.CommandContext(ctx, r.BinaryPath, args...)
	cmd.Dir = r.ProjectDir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("sqlmesh %s: %w\n%s", strings.Join(args, " "), err, out)
	}
	return nil
}
