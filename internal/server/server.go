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
}

func NewServer(stdout io.Writer) *Server {
	facade := facade.NewDummyFacade()
	router := NewRouter(facade)
	handler := createHandler(router)
	return &Server{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: handler,
		},
		stdout: stdout,
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
