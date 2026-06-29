package sources

import (
	"context"
	"net/http"

	"github.com/gear6io/pragmata/pkg/types/sourcetypes"
)

// Handler is the HTTP layer for source operations.
type Handler interface {
	CreateSource(w http.ResponseWriter, r *http.Request)
	ListSources(w http.ResponseWriter, r *http.Request)
	GetSource(w http.ResponseWriter, r *http.Request)
}

// Module is the business-logic layer for sources.
type Module interface {
	CreateSource(ctx context.Context, src *sourcetypes.Source) (*sourcetypes.Source, error)
	ListSources(ctx context.Context) ([]sourcetypes.Source, error)
	GetSource(ctx context.Context, name string) (*sourcetypes.Source, error)
}
