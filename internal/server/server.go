// Package server wires gin routes to the pipes Handler.
package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/gear6io/pragmata/pkg/modules/pipes"
	"github.com/gear6io/pragmata/pkg/sqlstore"
)

// Server is the HTTP API server.
type Server struct {
	router  *gin.Engine
	handler pipes.Handler
	store   sqlstore.SQLStore
}

// New builds a gin router with all pipe routes and auth middleware wired up.
func New(handler pipes.Handler, store sqlstore.SQLStore) *Server {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	s := &Server{router: r, handler: handler, store: store}

	v0 := r.Group("/v0")
	v0.Use(s.bearerAuth())

	// Pipe CRUD
	v0.POST("/pipes", s.wrap(handler.CreatePipe))
	v0.GET("/pipes", s.wrap(handler.ListPipes))

	// Named-pipe routes: inject pipe name into request context.
	named := v0.Group("/pipes/:name")
	named.Use(s.injectPipeName())
	named.GET("", s.wrap(handler.ExecutePipe))    // GET /v0/pipes/:name   → execute
	named.GET("/meta", s.wrap(handler.GetPipe))    // GET /v0/pipes/:name/meta → metadata
	named.PUT("", s.wrap(handler.UpdatePipe))
	named.DELETE("", s.wrap(handler.DeletePipe))

	return s
}

// Start listens on addr (e.g. ":7181") until ctx is cancelled.
func (s *Server) Start(ctx context.Context, addr string) error {
	srv := &http.Server{Addr: addr, Handler: s.router}
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

// wrap converts a standard http.HandlerFunc to a gin.HandlerFunc.
func (s *Server) wrap(h http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(c.Writer, c.Request)
	}
}

// injectPipeName reads the :name param from gin context and stores it in the
// request context so standard-library handlers can retrieve it via pipes.PipeName.
func (s *Server) injectPipeName() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := pipes.WithPipeName(c.Request, c.Param("name"))
		c.Request = r
		c.Next()
	}
}

// bearerAuth validates the Bearer token from the Authorization header or ?token= param.
func (s *Server) bearerAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := tokenFromRequest(c.Request)
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}
		_, err := s.store.GetTokenByValue(c.Request.Context(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Next()
	}
}

func tokenFromRequest(r *http.Request) string {
	// Authorization: Bearer <token>
	if auth := r.Header.Get("Authorization"); strings.HasPrefix(auth, "Bearer ") {
		return strings.TrimPrefix(auth, "Bearer ")
	}
	// ?token=<token>
	return r.URL.Query().Get("token")
}

// Handler returns the underlying http.Handler for testing.
func (s *Server) Handler() http.Handler {
	return s.router
}

// Addr formats host:port from config values.
func Addr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
