package http

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Server is a struct that encapsulates an HTTP server. It provides a simple interface to start and stop the server.
type Server struct {
	httpServer *http.Server
}

// NewHTTPServer creates a new HTTP server instance with the provided port and handler.
// The returned Server struct encapsulates the underlying http.Server and provides a simple
// interface to start and stop the server.
func NewHTTPServer(port int, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			Handler:      handler,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 20 * time.Second,
		},
	}
}

// Start runs the HTTP server in a separate goroutine. It returns immediately, and the server
// will continue running until it is stopped by calling the Stop method.
func (s *Server) Start() {
	go func() {
		// ListenAndServe always returns a non-nil error. After [Server.Shutdown] or [Server.Close],
		// the returned error is [ErrServerClosed].
		_ = s.httpServer.ListenAndServe()
	}()
}

// Stop gracefully shuts down the HTTP server. It waits for up to the specified
// context timeout for in-flight requests to complete before returning.
func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
