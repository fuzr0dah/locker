package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func createHandler(router *router) http.Handler {
	r := chi.NewRouter()
	r.Get("/status", router.handleStatus)
	r.Get("/secret/{id}", router.handleGetSecretByID)
	r.Post("/secret", router.handleCreateSecret)
	r.Put("/secret/{id}", router.handleUpdateSecret)
	return r
}
