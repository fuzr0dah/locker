package server

import (
	"net/http"

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

func (router *router) handleGetSecretByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	secret, err := router.facade.GetSecretById(r.Context(), idStr)
	if err != nil {
		respondWithError(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, secret)
}

func (router *router) handleUpdateSecret(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	var req api.UpdateSecretRequest
	if err := decodeRequest(r, &req); err != nil {
		badRequest(w, r, err.Error())
		return
	}

	secret, err := router.facade.UpdateSecret(r.Context(), idStr, req.Name, req.Value)
	if err != nil {
		respondWithError(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, secret)
}
