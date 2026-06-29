package implpipes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gear6io/pragmata/pkg/http/render"
	"github.com/gear6io/pragmata/pkg/modules/pipes"
	"github.com/gear6io/pragmata/pkg/pipevisitor"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

type handler struct {
	module pipes.Module
}

// NewHandler wraps a Module as an HTTP handler.
func NewHandler(mod pipes.Module) pipes.Handler {
	return &handler{module: mod}
}

func (h *handler) CreatePipe(w http.ResponseWriter, r *http.Request) {
	var body pipetypes.PostablePipe
	if !decodeJSON(w, r, &body) {
		return
	}

	created, err := h.module.CreatePipe(r.Context(), &body)
	if err != nil {
		render.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	render.Success(w, http.StatusCreated, created)
}

func (h *handler) ListPipes(w http.ResponseWriter, r *http.Request) {
	list, err := h.module.ListPipes(r.Context())
	if err != nil {
		render.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	if list == nil {
		list = []*pipetypes.GettablePipe{}
	}
	render.Success(w, http.StatusOK, list)
}

func (h *handler) GetPipe(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	pipe, err := h.module.GetPipe(r.Context(), name)
	if err != nil {
		render.Error(w, http.StatusNotFound, err.Error())
		return
	}
	render.Success(w, http.StatusOK, pipe)
}

func (h *handler) UpdatePipe(w http.ResponseWriter, r *http.Request) {
	var body pipetypes.PostablePipe
	if !decodeJSON(w, r, &body) {
		return
	}
	exec, err := pipevisitor.Visit(body.Content, pipevisitor.PipeVisitorOpts{})
	if err != nil {
		render.Error(w, http.StatusBadRequest, "parse error: "+err.Error())
		return
	}
	updated, err := h.module.UpdatePipe(r.Context(), exec)
	if err != nil {
		render.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	render.Success(w, http.StatusOK, updated)
}

func (h *handler) DeletePipe(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	if err := h.module.DeletePipe(r.Context(), name); err != nil {
		render.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func decodeJSON(w http.ResponseWriter, r *http.Request, v any) bool {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		render.Error(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return false
	}
	return true
}
