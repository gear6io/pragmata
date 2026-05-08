package implpipes

import (
	"encoding/json"
	"net/http"

	"github.com/gear6io/pragmata/pkg/modules/pipes"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

type handler struct {
	module pipes.Module
}

// NewHandler wraps a Module as an HTTP handler.
func NewHandler(mod pipes.Module) pipes.Handler {
	return &handler{module: mod}
}

// --- CRUD handlers ---

func (h *handler) CreatePipe(w http.ResponseWriter, r *http.Request) {
	var pipe pipetypes.Pipe
	if !decodeJSON(w, r, &pipe) {
		return
	}
	created, err := h.module.CreatePipe(r.Context(), &pipe)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

func (h *handler) ListPipes(w http.ResponseWriter, r *http.Request) {
	list, err := h.module.ListPipes(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if list == nil {
		list = []*pipetypes.Pipe{}
	}
	writeJSON(w, http.StatusOK, map[string]any{"pipes": list})
}

func (h *handler) GetPipe(w http.ResponseWriter, r *http.Request) {
	name := pipes.PipeName(r)
	pipe, err := h.module.GetPipe(r.Context(), name)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, pipe)
}

func (h *handler) UpdatePipe(w http.ResponseWriter, r *http.Request) {
	name := pipes.PipeName(r)
	var pipe pipetypes.Pipe
	if !decodeJSON(w, r, &pipe) {
		return
	}
	updated, err := h.module.UpdatePipe(r.Context(), name, &pipe)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

func (h *handler) DeletePipe(w http.ResponseWriter, r *http.Request) {
	name := pipes.PipeName(r)
	if err := h.module.DeletePipe(r.Context(), name); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// --- Execute handler ---

func (h *handler) ExecutePipe(w http.ResponseWriter, r *http.Request) {
	name := pipes.PipeName(r)

	// Collect URL query params as the template substitution map.
	params := make(map[string]string, len(r.URL.Query()))
	for k, vs := range r.URL.Query() {
		if len(vs) > 0 {
			params[k] = vs[0]
		}
	}

	result, err := h.module.ExecutePipe(r.Context(), name, params)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, result)
}

// --- helpers ---

func decodeJSON(w http.ResponseWriter, r *http.Request, v any) bool {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return false
	}
	return true
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
