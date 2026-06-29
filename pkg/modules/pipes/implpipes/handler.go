package implpipes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gear6io/pragmata/pkg/errors"
	"github.com/gear6io/pragmata/pkg/http/render"
	"github.com/gear6io/pragmata/pkg/modules/pipes"
	"github.com/gear6io/pragmata/pkg/pipevisitor"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

var (
	CodePipeNotFound       = errors.MustNewCode("pipe_not_found")
	CodeInvalidPipeContent = errors.MustNewCode("invalid_pipe_content")
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
		render.ErrorFrom(w, errors.WrapInternalf(err, errors.CodeInternal, "create pipe"))
		return
	}
	render.Success(w, http.StatusCreated, created)
}

func (h *handler) ListPipes(w http.ResponseWriter, r *http.Request) {
	list, err := h.module.ListPipes(r.Context())
	if err != nil {
		render.ErrorFrom(w, errors.WrapInternalf(err, errors.CodeInternal, "list pipes"))
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
		render.ErrorFrom(w, errors.WrapNotFoundf(err, CodePipeNotFound, "pipe %q not found", name))
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
		render.ErrorFrom(w, errors.WrapInvalidInputf(err, CodeInvalidPipeContent, "parse error"))
		return
	}
	updated, err := h.module.UpdatePipe(r.Context(), exec)
	if err != nil {
		render.ErrorFrom(w, errors.WrapInternalf(err, errors.CodeInternal, "update pipe"))
		return
	}
	render.Success(w, http.StatusOK, updated)
}

func (h *handler) DeletePipe(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	if err := h.module.DeletePipe(r.Context(), name); err != nil {
		render.ErrorFrom(w, errors.WrapInternalf(err, errors.CodeInternal, "delete pipe %q", name))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) ExecutePipe(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	params := make(map[string]string)
	for k, vals := range r.URL.Query() {
		if len(vals) > 0 {
			params[k] = vals[0]
		}
	}
	result, err := h.module.ExecutePipe(r.Context(), name, params)
	if err != nil {
		render.ErrorFrom(w, errors.WrapInternalf(err, errors.CodeInternal, "execute pipe %q", name))
		return
	}
	render.Success(w, http.StatusOK, result)
}

func decodeJSON(w http.ResponseWriter, r *http.Request, v any) bool {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		render.ErrorFrom(w, errors.WrapInvalidInputf(err, errors.CodeInvalidInput, "invalid JSON"))
		return false
	}
	return true
}
