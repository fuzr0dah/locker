package server

import (
	"log/slog"

	"github.com/fuzr0dah/locker/internal/api"
	"github.com/go-chi/chi/v5"
)

func createHandler(router *router, logger *slog.Logger) *chi.Mux {
	r := chi.NewRouter()

	r.Use(loggerMiddleware(logger))

	r.MethodFunc(api.Secrets.Create.Method, api.Secrets.Create.Path, router.handleCreateSecret)
	r.MethodFunc(api.Secrets.Get.Method, api.Secrets.Get.Path, router.handleGetSecretByID)
	r.MethodFunc(api.Secrets.Update.Method, api.Secrets.Update.Path, router.handleUpdateSecret)
	r.MethodFunc(api.Secrets.Delete.Method, api.Secrets.Delete.Path, router.handleDeleteSecret)

	r.MethodFunc(api.Secrets.Versions.Get.Method, api.Secrets.Versions.Get.Path, router.handleGetSecretVersion)

	return r
}
