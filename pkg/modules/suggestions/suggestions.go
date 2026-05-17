package suggestions

import (
	"context"
	"net/http"

	"github.com/gear6io/pragmata/pkg/types/suggestiontypes"
)

// Handler is the HTTP layer for suggestion operations.
type Handler interface {
	GetSuggestions(w http.ResponseWriter, r *http.Request)
}

// Module is the business-logic layer for suggestions.
type Module interface {
	GetSuggestions(ctx context.Context, req suggestiontypes.SuggestionRequest) (*suggestiontypes.SuggestionResponse, error)
}
