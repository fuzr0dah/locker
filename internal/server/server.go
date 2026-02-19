package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/fuzr0dah/locker/internal/crypto"
	"github.com/fuzr0dah/locker/internal/db/migrations"
	"github.com/fuzr0dah/locker/internal/facade"
	"github.com/fuzr0dah/locker/internal/service"
	"github.com/fuzr0dah/locker/internal/storage"
	"github.com/fuzr0dah/locker/internal/storage/sqlite"
)

type Server struct {
	httpServer *http.Server
	logger     *slog.Logger
	storage    storage.SecretStorage
}

func NewServer(logger *slog.Logger) (*Server, error) {
	if logger == nil {
		return nil, errors.New("logger is required")
	}
	logger.Info("initializing server", "addr", ":8080")
	logger.Info("master key generated", "key", crypto.GenerateMasterKey())

	db, err := sqlite.OpenDB("")
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	if err := migrations.RunMigrations(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("run migrations: %w", err)
	}

	storage := sqlite.NewStorage(db)
	svc := service.NewSecretsService(storage)

	facadeLogger := logger.With("component", "facade")
	facade := facade.NewFacade(svc, facadeLogger)

	handlerLogger := logger.With("component", "handler")
	router := newRouter(facade)
	handler := createHandler(router, handlerLogger)

	return &Server{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: handler,
		},
		logger:  logger,
		storage: storage,
	}, nil
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("http shutdown: %w", err)
	}
	if err := s.storage.Close(); err != nil {
		return fmt.Errorf("storage close: %w", err)
	}
	return nil
}
