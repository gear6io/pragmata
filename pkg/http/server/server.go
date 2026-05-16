// Package server wires gorilla/mux routes to the pipes Handler.
package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/swaggest/openapi-go/openapi3"

	"github.com/gear6io/pragmata/pkg/http/handler"
	"github.com/gear6io/pragmata/pkg/http/render"
	"github.com/gear6io/pragmata/pkg/modules/pipes"
	"github.com/gear6io/pragmata/pkg/sqlstore"
	"github.com/gear6io/pragmata/pkg/types/pipetypes"
)

// Provider groups all module handlers the HTTP server routes to.
// Add a new field here when a second module lands — keeps New() signature stable.
type Provider struct {
	Pipes pipes.Handler
}

// bearerScheme is the single declared security scheme applied to all v0 routes.
var bearerScheme = []handler.OpenAPISecurityScheme{{Name: "BearerAuth"}}

// Server is the HTTP API server.
type Server struct {
	router   *mux.Router
	provider *Provider
	store    sqlstore.SQLStore
	addr     string
	openapi  *handler.OpenAPICollector
}

// New builds a mux router with all routes, bearer auth, and OpenAPI collection wired up.
func New(p *Provider, store sqlstore.SQLStore, addr string) *Server {
	reflector := openapi3.NewReflector()
	reflector.Spec.WithInfo(*(&openapi3.Info{}).
		WithTitle("Pragmata API").
		WithVersion("v0"))
	reflector.SpecSchema().SetHTTPBearerTokenSecurity("BearerAuth", "", "API bearer token")

	oac := handler.NewOpenAPICollector(reflector)

	r := mux.NewRouter()
	r.Use(recoveryMiddleware)

	s := &Server{router: r, provider: p, store: store, addr: addr, openapi: oac}

	v0 := r.PathPrefix("/v0").Subrouter()
	v0.Use(s.bearerAuth)
	v0.Use(s.injectPipeName) // no-op on routes without {name}

	h := p.Pipes

	v0.Handle("/pipes", handler.New(h.CreatePipe, handler.OpenAPIDef{
		ID:                "createPipe",
		Tags:              []string{"pipes"},
		Summary:           "Create a pipe",
		Request:           new(pipetypes.Pipe),
		Response:          new(pipetypes.Pipe),
		SuccessStatusCode: http.StatusCreated,
		ErrorStatusCodes:  []int{http.StatusBadRequest, http.StatusInternalServerError},
		SecuritySchemes:   bearerScheme,
	})).Methods("POST")

	v0.Handle("/pipes", handler.New(h.ListPipes, handler.OpenAPIDef{
		ID:              "listPipes",
		Tags:            []string{"pipes"},
		Summary:         "List all pipes",
		Response:        []*pipetypes.Pipe{}, // slice schema: Data is an array of Pipe
		ErrorStatusCodes: []int{http.StatusInternalServerError},
		SecuritySchemes: bearerScheme,
	})).Methods("GET")

	v0.Handle("/pipes/{name}/meta", handler.New(h.GetPipe, handler.OpenAPIDef{
		ID:               "getPipe",
		Tags:             []string{"pipes"},
		Summary:          "Get pipe metadata",
		Response:         new(pipetypes.Pipe),
		ErrorStatusCodes: []int{http.StatusNotFound},
		SecuritySchemes:  bearerScheme,
	})).Methods("GET")

	v0.Handle("/pipes/{name}", handler.New(h.UpdatePipe, handler.OpenAPIDef{
		ID:               "updatePipe",
		Tags:             []string{"pipes"},
		Summary:          "Update a pipe",
		Request:          new(pipetypes.Pipe),
		Response:         new(pipetypes.Pipe),
		ErrorStatusCodes: []int{http.StatusBadRequest, http.StatusNotFound, http.StatusInternalServerError},
		SecuritySchemes:  bearerScheme,
	})).Methods("PUT")

	v0.Handle("/pipes/{name}", handler.New(h.DeletePipe, handler.OpenAPIDef{
		ID:                "deletePipe",
		Tags:              []string{"pipes"},
		Summary:           "Delete a pipe",
		SuccessStatusCode: http.StatusNoContent,
		ErrorStatusCodes:  []int{http.StatusNotFound, http.StatusInternalServerError},
		SecuritySchemes:   bearerScheme,
	})).Methods("DELETE")

	// Walk all registered routes once to populate the OpenAPI collector.
	if err := r.Walk(oac.Walker); err != nil {
		panic("openapi: walk routes: " + err.Error())
	}

	// Serve the spec at runtime without auth.
	r.Handle("/openapi.yaml", http.HandlerFunc(s.serveOpenAPIYAML))

	return s
}

// Start listens on the configured addr until ctx is cancelled.
func (s *Server) Start(ctx context.Context) error {
	srv := &http.Server{Addr: s.addr, Handler: s.router}
	errCh := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()
	select {
	case <-ctx.Done():
		return srv.Shutdown(context.Background())
	case err := <-errCh:
		return err
	}
}

// Stop is a no-op: Start handles graceful shutdown via ctx cancellation.
func (s *Server) Stop(_ context.Context) error { return nil }

// Handler returns the underlying http.Handler for testing.
func (s *Server) Handler() http.Handler { return s.router }

// OpenAPISpec returns the accumulated spec as YAML bytes. Used by generate command.
func (s *Server) OpenAPISpec() ([]byte, error) { return s.openapi.MarshalYAML() }

// Addr formats host:port from config values.
func Addr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func (s *Server) serveOpenAPIYAML(w http.ResponseWriter, _ *http.Request) {
	data, err := s.openapi.MarshalYAML()
	if err != nil {
		render.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/x-yaml")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func (s *Server) bearerAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := tokenFromRequest(r)
		if token == "" {
			render.Error(w, http.StatusUnauthorized, "missing token")
			return
		}
		if _, err := s.store.GetTokenByValue(r.Context(), token); err != nil {
			render.Error(w, http.StatusUnauthorized, "invalid token")
			return
		}
		next.ServeHTTP(w, r)
	})
}

// injectPipeName reads the {name} path variable and stores it in the request context
// so stdlib handlers can retrieve it via pipes.PipeName. No-op on routes without {name}.
func (s *Server) injectPipeName(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if name := mux.Vars(r)["name"]; name != "" {
			r = pipes.WithPipeName(r, name)
		}
		next.ServeHTTP(w, r)
	})
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				render.Error(w, http.StatusInternalServerError, "internal server error")
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func tokenFromRequest(r *http.Request) string {
	if auth := r.Header.Get("Authorization"); strings.HasPrefix(auth, "Bearer ") {
		return strings.TrimPrefix(auth, "Bearer ")
	}
	return r.URL.Query().Get("token")
}
