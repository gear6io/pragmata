package implpipes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gear6io/pragmata/pkg/http/render"
	"github.com/gear6io/pragmata/pkg/modules/pipes"
	"github.com/gear6io/pragmata/pkg/pipevisitor"
	"github.com/gear6io/pragmata/pkg/types"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
	"github.com/gear6io/pragmata/pkg/valuer"
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
	pipe, err := pipevisitor.Visit("", body.Content)
	if err != nil {
		render.Error(w, http.StatusBadRequest, "parse error: "+err.Error())
		return
	}

	// create storable flavor
	storable := pipetypes.StorablePipe{
		Identifiable: types.Identifiable{
			ID: valuer.GenerateUUID(),
		},
		Pipe: *pipe,
		TimeAuditable: types.TimeAuditable{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	created, err := h.module.CreatePipe(r.Context(), &storable)
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
		list = []*pipetypes.Pipe{}
	}
	render.Success(w, http.StatusOK, list)
}

func (h *handler) GetPipe(w http.ResponseWriter, r *http.Request) {
	name := pipes.PipeName(r)
	pipe, err := h.module.GetPipe(r.Context(), name)
	if err != nil {
		render.Error(w, http.StatusNotFound, err.Error())
		return
	}
	render.Success(w, http.StatusOK, pipe)
}

func (h *handler) UpdatePipe(w http.ResponseWriter, r *http.Request) {
	name := pipes.PipeName(r)
	var body pipetypes.PostablePipe
	if !decodeJSON(w, r, &body) {
		return
	}
	pipe, err := pipevisitor.Visit(name, body.Content)
	if err != nil {
		render.Error(w, http.StatusBadRequest, "parse error: "+err.Error())
		return
	}
	// create storable flavor
	storable := &pipetypes.StorablePipe{
		Pipe: *pipe,
		TimeAuditable: types.TimeAuditable{
			UpdatedAt: time.Now(),
		},
	}
	updated, err := h.module.UpdatePipe(r.Context(), storable)
	if err != nil {
		render.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	render.Success(w, http.StatusOK, updated)
}

func (h *handler) DeletePipe(w http.ResponseWriter, r *http.Request) {
	name := pipes.PipeName(r)
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
