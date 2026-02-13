package server

import (
	"net/http"

	"github.com/fuzr0dah/locker/internal/facade"

	"github.com/go-chi/render"
)

type router struct {
	facade facade.SecretsFacade
}

func newRouter(f facade.SecretsFacade) *router {
	return &router{facade: f}
}

func (router *router) handleStatus(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, `{"status":"ok"}`)
}
