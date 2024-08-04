package http

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewHTTPServer(port int, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d",port),
			Handler: handler,
		},
	}
}

func (s *Server) Start() {
	go func() {
		// ListenAndServe always returns a non-nil error. After [Server.Shutdown] or [Server.Close],
		// the returned error is [ErrServerClosed].
		_ = s.httpServer.ListenAndServe()
	}()
}

func (s *Server) Stop(ctx context.Context) error{
	return s.httpServer.Shutdown(ctx)
}