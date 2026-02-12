package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func createHandler(router *Router) http.Handler {
	r := chi.NewRouter()
	r.Get("/status", router.handleStatus)
	return r
}
