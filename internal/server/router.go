package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/fuzr0dah/locker/internal/dto"
	"github.com/fuzr0dah/locker/internal/facade"
	"github.com/fuzr0dah/locker/internal/secrets"

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
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "invalid id"})
		return
	}

	secret, err := router.facade.GetSecretById(r.Context(), id)
	if err != nil {
		if errors.Is(err, secrets.ErrSecretNotFound) {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"error": "secret not found"})
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	resp := dto.GetSecretResponse{
		ID:        secret.ID,
		Name:      secret.Name,
		Value:     string(secret.Value),
		Version:   secret.Version,
		CreatedAt: secret.CreatedAt,
		UpdatedAt: secret.UpdatedAt,
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp)
}

func (router *router) handleCreateSecret(w http.ResponseWriter, r *http.Request) {
	var data dto.CreateSecretRequest
	if err := render.Bind(r, &data); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	secret, err := router.facade.CreateSecret(r.Context(), data.Name, data.Value)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	resp := dto.CreateSecretResponse{
		ID:        secret.ID,
		Name:      secret.Name,
		Value:     string(secret.Value),
		Version:   secret.Version,
		CreatedAt: secret.CreatedAt,
		UpdatedAt: secret.UpdatedAt,
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp)
}
