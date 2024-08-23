package metrics

import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Server is a struct that encapsulates an HTTP server. It provides a simple interface to start and stop the server.
type Server struct {
	httpServer *http.Server
}

// NewMetricsServer creates a new HTTP server instance specifically for serving Prometheus metrics.
// It sets up a new ServeMux and registers the Prometheus handler at the "/metrics" endpoint.
// The returned Server struct encapsulates the underlying http.Server and provides a simple
// interface to start and stop the server.
func NewMetricsServer() *Server {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	return &Server{
		httpServer: &http.Server{
			Addr:         ":2112",
			Handler:      mux,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
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
