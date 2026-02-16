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
	"github.com/fuzr0dah/locker/internal/secrets"
)

type Server struct {
	httpServer *http.Server
	logger     *slog.Logger
	storage    secrets.Storage
}

func NewServer(logger *slog.Logger) (*Server, error) {
	if logger == nil {
		return nil, errors.New("logger is required")
	}
	logger.Info("initializing server", "addr", ":8080")
	logger.Info("master key generated", "key", crypto.GenerateMasterKey())

	db, err := secrets.OpenDB()
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	if err := migrations.RunMigrations(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("run migrations: %w", err)
	}

	storage := secrets.NewStorage(db)
	service := secrets.NewService(storage)

	facadeLogger := logger.With("component", "facade")
	facade := facade.NewFacade(service, facadeLogger)

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
