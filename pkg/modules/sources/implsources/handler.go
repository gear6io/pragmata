package implsources

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gear6io/pragmata/pkg/errors"
	"github.com/gear6io/pragmata/pkg/http/render"
	"github.com/gear6io/pragmata/pkg/modules/sources"
	"github.com/gear6io/pragmata/pkg/types/sourcetypes"
)

var CodeSourceNotFound = errors.MustNewCode("source_not_found")

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
		render.ErrorFrom(w, errors.WrapInvalidInputf(err, errors.CodeInvalidInput, "invalid JSON"))
		return
	}
	created, err := h.module.CreateSource(r.Context(), &src)
	if err != nil {
		render.ErrorFrom(w, errors.WrapInternalf(err, errors.CodeInternal, "create source"))
		return
	}
	render.Success(w, http.StatusCreated, created)
}

func (h *handler) ListSources(w http.ResponseWriter, r *http.Request) {
	list, err := h.module.ListSources(r.Context())
	if err != nil {
		render.ErrorFrom(w, errors.WrapInternalf(err, errors.CodeInternal, "list sources"))
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
		render.ErrorFrom(w, errors.WrapInternalf(err, errors.CodeInternal, "get source %q", name))
		return
	}
	if src == nil {
		render.ErrorFrom(w, errors.NewNotFoundf(CodeSourceNotFound, "source %q not found", name))
		return
	}
	render.Success(w, http.StatusOK, src)
}
