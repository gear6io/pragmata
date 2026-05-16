package handler

import (
	"net/http"

	openapigo "github.com/swaggest/openapi-go"

	"github.com/gear6io/pragmata/pkg/http/render"
)

// OpenAPIDef holds OpenAPI metadata for a single HTTP operation.
// Port of SigNoz's pkg/http/handler.OpenAPIDef, without authz/audit fields.
type OpenAPIDef struct {
	ID                string
	Tags              []string
	Summary           string
	Description       string
	Request           any // pointer to request body struct, or nil
	RequestQuery      any // pointer to query-params struct, or nil
	Response          any // pointer to success response struct, or nil
	SuccessStatusCode int // defaults to 200 when zero
	ErrorStatusCodes  []int
	SecuritySchemes   []OpenAPISecurityScheme
}

// OpenAPISecurityScheme names a declared security scheme and its optional scopes.
type OpenAPISecurityScheme struct {
	Name   string
	Scopes []string
}

// Handler is an http.Handler that can also describe itself to an OpenAPI collector.
type Handler interface {
	http.Handler
	ServeOpenAPI(oc openapigo.OperationContext)
}

type handler struct {
	fn  http.HandlerFunc
	def OpenAPIDef
}

// New wraps fn with OpenAPI metadata from def.
func New(fn http.HandlerFunc, def OpenAPIDef) Handler {
	return &handler{fn: fn, def: def}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.fn(w, r)
}

func (h *handler) ServeOpenAPI(oc openapigo.OperationContext) {
	oc.SetID(h.def.ID)
	oc.SetTags(h.def.Tags...)
	oc.SetSummary(h.def.Summary)
	oc.SetDescription(h.def.Description)

	for _, s := range h.def.SecuritySchemes {
		oc.AddSecurity(s.Name, s.Scopes...)
	}
	if h.def.Request != nil {
		oc.AddReqStructure(h.def.Request)
	}
	if h.def.RequestQuery != nil {
		oc.AddReqStructure(h.def.RequestQuery)
	}

	successCode := h.def.SuccessStatusCode
	if successCode == 0 {
		successCode = http.StatusOK
	}
	if h.def.Response != nil {
		oc.AddRespStructure(h.def.Response, openapigo.WithHTTPStatus(successCode))
	} else {
		oc.AddRespStructure(nil, openapigo.WithHTTPStatus(successCode))
	}
	// render.ErrorResponse matches the actual shape written by render.Error.
	for _, code := range h.def.ErrorStatusCodes {
		oc.AddRespStructure(render.ErrorResponse{}, openapigo.WithHTTPStatus(code))
	}
}
