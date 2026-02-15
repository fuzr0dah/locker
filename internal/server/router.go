package server

import (
	"net/http"
	"strconv"

	"github.com/fuzr0dah/locker/internal/api"
	"github.com/fuzr0dah/locker/internal/facade"

	"github.com/go-chi/chi/v5"
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

func (router *router) handleGetSecretByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		badRequest(w, r, "invalid id")
		return
	}

	secret, err := router.facade.GetSecretById(r.Context(), id)
	if err != nil {
		respondWithError(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, secret)
}

func (router *router) handleCreateSecret(w http.ResponseWriter, r *http.Request) {
	var req api.CreateSecretRequest
	if err := decodeRequest(r, &req); err != nil {
		badRequest(w, r, err.Error())
		return
	}

	secret, err := router.facade.CreateSecret(r.Context(), req.Name, req.Value)
	if err != nil {
		respondWithError(w, r, err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, secret)
}
