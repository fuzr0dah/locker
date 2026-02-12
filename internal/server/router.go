package server

import (
	"net/http"

	"github.com/fuzr0dah/locker/internal/facade"

	"github.com/go-chi/render"
)

type Router struct {
	Facade facade.DummyFacade
}

func NewRouter(facade facade.DummyFacade) *Router {
	return &Router{Facade: facade}
}

func (router *Router) handleStatus(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, `{"status":"ok"}`)
}
