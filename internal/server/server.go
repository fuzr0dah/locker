package server

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/fuzr0dah/locker/internal/db/migrations"
	"github.com/fuzr0dah/locker/internal/facade"
	"github.com/fuzr0dah/locker/internal/secrets"
)

type Server struct {
	httpServer *http.Server
	stdout     io.Writer
	storage    secrets.Storage
}

func NewServer(stdout io.Writer) (*Server, error) {
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
	facade := facade.NewFacade(service)
	router := newRouter(facade)
	handler := createHandler(router)

	return &Server{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: handler,
		},
		stdout:  stdout,
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
