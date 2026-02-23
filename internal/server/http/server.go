package http

import (
	"log/slog"
	"net/http"

	"github.com/fuzr0dah/locker/internal/application/facade"
)

type server struct {
	httpServer *http.Server
	logger     *slog.Logger
}

func NewServer(facade facade.SecretsFacade, logger *slog.Logger) (*server, error) {
	logger.Info("initializing server", "addr", ":8080")

	handlerLogger := logger.With("component", "handler")
	router := newRouter(facade)
	handler := createHandler(router, handlerLogger)

	return &server{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: handler,
		},
		logger: logger,
	}, nil
}

func (s *server) Start() error {
	return s.httpServer.ListenAndServe()
}
