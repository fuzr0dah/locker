package http

import (
	"net/http"
	"strconv"

	"github.com/fuzr0dah/locker/internal/api"
	"github.com/fuzr0dah/locker/internal/application/facade"
	"github.com/fuzr0dah/locker/internal/domain/secrets"

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
		respondWithError(w, r, err)
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
	idParam := chi.URLParam(r, "id")
	if !secrets.IsValid(idParam) {
		badRequest(w, r, "invalid secret id format")
		return
	}

	secret, err := router.facade.GetSecretById(r.Context(), idParam)
	if err != nil {
		respondWithError(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, secret)
}

func (router *router) handleUpdateSecret(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if !secrets.IsValid(idParam) {
		badRequest(w, r, "invalid secret id format")
		return
	}

	var req api.UpdateSecretRequest
	if err := decodeRequest(r, &req); err != nil {
		respondWithError(w, r, err)
		return
	}

	secret, err := router.facade.UpdateSecret(r.Context(), idParam, req.Name, req.Value)
	if err != nil {
		respondWithError(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, secret)
}

func (router *router) handleDeleteSecret(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if !secrets.IsValid(idParam) {
		badRequest(w, r, "invalid secret id format")
		return
	}

	if err := router.facade.DeleteSecret(r.Context(), idParam); err != nil {
		respondWithError(w, r, err)
		return
	}

	render.Status(r, http.StatusNoContent)
}

func (router *router) handleGetSecretVersion(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if !secrets.IsValid(idParam) {
		badRequest(w, r, "invalid secret id format")
		return
	}

	versionStr := chi.URLParam(r, "version")
	version, err := strconv.Atoi(versionStr)
	if err != nil || version <= 0 {
		badRequest(w, r, "version must be a positive integer")
		return
	}

	secret, err := router.facade.GetSecretVersion(r.Context(), idParam, version)
	if err != nil {
		respondWithError(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, secret)
}
