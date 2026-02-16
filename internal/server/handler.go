package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func createHandler(router *router, logger *slog.Logger) http.Handler {
	r := chi.NewRouter()

	r.Use(loggerMiddleware(logger))

	r.Get("/status", router.handleStatus)
	r.Get("/secret/{id}", router.handleGetSecretByID)
	r.Post("/secret", router.handleCreateSecret)
	r.Put("/secret/{id}", router.handleUpdateSecret)

	return r
}
