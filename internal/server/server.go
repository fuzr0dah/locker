package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/fuzr0dah/locker/internal/application/facade"
)

type Server struct {
	httpServer *http.Server
	logger     *slog.Logger
}

func NewServer(facade facade.SecretsFacade, logger *slog.Logger) (*Server, error) {
	logger.Info("initializing server", "addr", ":8080")

	handlerLogger := logger.With("component", "handler")
	router := newRouter(facade)
	handler := createHandler(router, handlerLogger)

	return &Server{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: handler,
		},
		logger: logger,
	}, nil
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("http shutdown: %w", err)
	}
	return nil
}
