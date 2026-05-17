package implsuggestions

import (
	"encoding/json"
	"net/http"

	"github.com/gear6io/pragmata/pkg/http/render"
	"github.com/gear6io/pragmata/pkg/modules/suggestions"
	"github.com/gear6io/pragmata/pkg/types/suggestiontypes"
)

type handler struct {
	module suggestions.Module
}

// NewHandler wraps a Module as an HTTP handler.
func NewHandler(mod suggestions.Module) suggestions.Handler {
	return &handler{module: mod}
}

func (h *handler) GetSuggestions(w http.ResponseWriter, r *http.Request) {
	var req suggestiontypes.SuggestionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		render.Error(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}

	if req.MatchingType == "" {
		req.MatchingType = suggestiontypes.MatchingTypeFuzzy
	}

	resp, err := h.module.GetSuggestions(r.Context(), req)
	if err != nil {
		render.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	render.Success(w, http.StatusOK, resp)
}
