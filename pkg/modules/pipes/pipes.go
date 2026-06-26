package pipes

import (
	"context"
	"net/http"

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
	// ExecutePipe(w http.ResponseWriter, r *http.Request)
}

// Module is the business logic layer for pipe operations.
type Module interface {
	CreatePipe(ctx context.Context, pipe *pipetypes.PostablePipe) (*pipetypes.GettablePipe, error)
	GetPipe(ctx context.Context, name string) (*pipetypes.GettablePipe, error)
	ListPipes(ctx context.Context) ([]*pipetypes.GettablePipe, error)
	UpdatePipe(ctx context.Context, exec *pipetypes.ExecutablePipe) (*pipetypes.GettablePipe, error)
	DeletePipe(ctx context.Context, name string) error
	// ExecutePipe(ctx context.Context, name string, params map[string]string) (*pipetypes.ExecuteResult, error)
}
