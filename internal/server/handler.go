package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func createHandler(router *router) http.Handler {
	r := chi.NewRouter()
	r.Get("/status", router.handleStatus)
	r.Get("/secret", router.handleGetSecret)
	r.Post("/secret", router.handleCreateSecret)
	return r
}
