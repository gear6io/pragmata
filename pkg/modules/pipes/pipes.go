package pipes

import (
	"context"
	"net/http"

	"github.com/gear6io/pragmata/pkg/scheduler"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

// Handler is the HTTP layer for pipe operations.
type Handler interface {
	// CRUD
	CreatePipe(w http.ResponseWriter, r *http.Request)
	ListPipes(w http.ResponseWriter, r *http.Request)
	GetPipe(w http.ResponseWriter, r *http.Request)
	UpdatePipe(w http.ResponseWriter, r *http.Request)
	DeletePipe(w http.ResponseWriter, r *http.Request)
	// Execution
	ExecutePipe(w http.ResponseWriter, r *http.Request)
}

// Module is the business logic layer for pipe operations.
type Module interface {
	CreatePipe(ctx context.Context, pipe *pipetypes.Pipe) (*pipetypes.Pipe, error)
	GetPipe(ctx context.Context, name string) (*pipetypes.Pipe, error)
	ListPipes(ctx context.Context) ([]*pipetypes.Pipe, error)
	UpdatePipe(ctx context.Context, name string, pipe *pipetypes.Pipe) (*pipetypes.Pipe, error)
	DeletePipe(ctx context.Context, name string) error
	ExecutePipe(ctx context.Context, name string, params map[string]string) (*pipetypes.ExecuteResult, error)
}

// Scheduler is re-exported from pkg/scheduler for use as a dependency type in this module.
type Scheduler = scheduler.Scheduler

// contextKey is a package-scoped type for context values to avoid collisions.
type contextKey string

const pipeNameKey contextKey = "pipe_name"

// WithPipeName returns a copy of r with the pipe name stored in its context.
func WithPipeName(r *http.Request, name string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), pipeNameKey, name))
}

// PipeName retrieves the pipe name injected by the router middleware.
func PipeName(r *http.Request) string {
	name, _ := r.Context().Value(pipeNameKey).(string)
	return name
}
