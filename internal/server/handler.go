package server

import (
	"log/slog"
	"net/http"

	"github.com/fuzr0dah/locker/internal/api"
	"github.com/go-chi/chi/v5"
)

func createHandler(router *router, logger *slog.Logger) http.Handler {
	r := chi.NewRouter()

	r.Use(loggerMiddleware(logger))

	r.MethodFunc(api.CreateSecret.Method, api.CreateSecret.Path, router.handleCreateSecret)
	r.MethodFunc(api.GetSecret.Method, api.GetSecret.Path, router.handleGetSecretByID)
	r.MethodFunc(api.UpdateSecret.Method, api.UpdateSecret.Path, router.handleUpdateSecret)

	return r
}
