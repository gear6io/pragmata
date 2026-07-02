// Package sqlmesh provides Go wrappers around the SQLMesh CLI.
package sqlmesh

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gear6io/pragmata/pkg/errors"
	"github.com/gear6io/pragmata/pkg/sqlmeshbuilder"
	"github.com/gear6io/pragmata/pkg/types/executortypes"
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

// EnsureProject creates ProjectDir and writes a minimal config.yaml if absent.
// Call once on startup before PlanApply.
func (r *Runner) EnsureProject(clickhouseURL string) error {
	if err := os.MkdirAll(r.ProjectDir, 0o755); err != nil {
		return errors.WrapInternalf(err, errors.CodeInternal, "create sqlmesh project dir")
	}
	cfgPath := filepath.Join(r.ProjectDir, "config.yaml")
	if _, err := os.Stat(cfgPath); err == nil {
		return nil // already exists
	}
	u, err := url.Parse(clickhouseURL)
	if err != nil {
		return errors.WrapInvalidInputf(err, errors.CodeInvalidInput, "parse clickhouse url")
	}
	host := u.Hostname()
	port := u.Port()
	if port == "" {
		port = "9000"
	}
	username := u.User.Username()
	password, _ := u.User.Password()
	// ponytail: minimal config; extend when multi-gateway or audit table needed
	// database field omitted — rejected by sqlmesh clickhouse connector as extra input
	cfg := fmt.Sprintf(`gateways:
  default:
    connection:
      type: clickhouse
      host: %s
      port: %s
      username: %s
      password: %s

model_defaults:
  dialect: clickhouse
`, host, port, username, password)
	return os.WriteFile(cfgPath, []byte(cfg), 0o644)
}

// SyncModel writes (or updates) the SQLMesh .sql model file for pipe.
func (r *Runner) SyncModel(_ context.Context, pipe *pipetypes.Pipe) error {
	sql, err := sqlmeshbuilder.FromPipe(pipe)
	if err != nil {
		dest := filepath.Join(r.ProjectDir, "models", pipe.Name+".sql")
		_ = os.Remove(dest)
		return errors.WrapInternalf(err, errors.CodeInternal, "generate model for %q", pipe.Name)
	}
	modelsDir := filepath.Join(r.ProjectDir, "models")
	if err := os.MkdirAll(modelsDir, 0o755); err != nil {
		return errors.WrapInternalf(err, errors.CodeInternal, "create models dir")
	}
	dest := filepath.Join(modelsDir, pipe.Name+".sql")
	return os.WriteFile(dest, []byte(sql), 0o644)
}

// RemoveModel deletes the SQLMesh .sql model file for the given pipe name.
func (r *Runner) RemoveModel(_ context.Context, name string) error {
	dest := filepath.Join(r.ProjectDir, "models", name+".sql")
	if err := os.Remove(dest); err != nil && !errors.Is(err, os.ErrNotExist) {
		return errors.WrapInternalf(err, errors.CodeInternal, "remove model %q", dest)
	}
	return nil
}

// PlanApply runs `sqlmesh plan --auto-apply` in ProjectDir.
func (r *Runner) PlanApply(ctx context.Context) error {
	return r.run(ctx, "plan", "--auto-apply")
}

// Run executes `sqlmesh run --select-model name [--start S --end E]`.
// interval may be nil for a full run.
func (r *Runner) Run(ctx context.Context, modelName string, interval *executortypes.TimeInterval) error {
	args := []string{"run", "--select-model", modelName}
	if interval != nil {
		args = append(args,
			"--start", interval.Start.Format("2006-01-02"),
			"--end", interval.End.Format("2006-01-02"),
		)
	}
	return r.run(ctx, args...)
}

// BackfillIntervals returns the date intervals for incremental backfill.
// ponytail: returns nil (no intervals) → initial load is done by PlanApply for FULL models.
// Extend when Pipe gains BackfillStart + time_column for INCREMENTAL_BY_TIME_RANGE kind.
func (r *Runner) BackfillIntervals(_ *pipetypes.Pipe) []executortypes.TimeInterval {
	return nil
}

// run executes the sqlmesh binary with the given args inside ProjectDir.
// A non-zero exit code becomes an error that includes captured stdout+stderr.
func (r *Runner) run(ctx context.Context, args ...string) error {
	cmd := exec.CommandContext(ctx, r.BinaryPath, args...)
	cmd.Dir = r.ProjectDir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.WrapInternalf(err, errors.CodeInternal, "sqlmesh %s:\n%s", strings.Join(args, " "), out)
	}
	return nil
}
