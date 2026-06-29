package implsources

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gear6io/pragmata/pkg/http/render"
	"github.com/gear6io/pragmata/pkg/modules/sources"
	"github.com/gear6io/pragmata/pkg/types/sourcetypes"
)

type handler struct {
	module sources.Module
}

// NewHandler wraps a Module as an HTTP handler.
func NewHandler(mod sources.Module) sources.Handler {
	return &handler{module: mod}
}

func (h *handler) CreateSource(w http.ResponseWriter, r *http.Request) {
	var src sourcetypes.Source
	if err := json.NewDecoder(r.Body).Decode(&src); err != nil {
		render.Error(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	created, err := h.module.CreateSource(r.Context(), &src)
	if err != nil {
		render.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	render.Success(w, http.StatusCreated, created)
}

func (h *handler) ListSources(w http.ResponseWriter, r *http.Request) {
	list, err := h.module.ListSources(r.Context())
	if err != nil {
		render.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	if list == nil {
		list = []sourcetypes.Source{}
	}
	render.Success(w, http.StatusOK, list)
}

func (h *handler) GetSource(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	src, err := h.module.GetSource(r.Context(), name)
	if err != nil {
		render.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	if src == nil {
		render.Error(w, http.StatusNotFound, "source not found")
		return
	}
	render.Success(w, http.StatusOK, src)
}
