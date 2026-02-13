package server

import (
	"context"
	"io"
	"net/http"

	"github.com/fuzr0dah/locker/internal/facade"
)

type Server struct {
	httpServer *http.Server
	stdout     io.Writer
	facade     facade.SecretsFacade
}

// NewServer creates a new server with the given facade
func NewServer(stdout io.Writer, f facade.SecretsFacade) *Server {
	router := newRouter(f)
	handler := createHandler(router)
	return &Server{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: handler,
		},
		stdout: stdout,
		facade: f,
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
