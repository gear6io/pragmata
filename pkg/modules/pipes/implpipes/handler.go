package implpipes

import (
	"net/http"

	"github.com/gear6io/pragmata/pkg/modules/pipes"
)

type handler struct {
	module pipes.Module
}

func NewHandler(mod pipes.Module) pipes.Handler {
	return &handler{module: mod}
}

func (h *handler) CreatePipe(w http.ResponseWriter, r *http.Request) {
}
func (h *handler) ListPipes(w http.ResponseWriter, r *http.Request) {
}
func (h *handler) GetPipe(w http.ResponseWriter, r *http.Request) {
}
func (h *handler) UpdatePipe(w http.ResponseWriter, r *http.Request) {
}
func (h *handler) DeletePipe(w http.ResponseWriter, r *http.Request) {
}
