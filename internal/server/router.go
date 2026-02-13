package server

import (
	"net/http"

	"github.com/fuzr0dah/locker/internal/dto"
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

func (router *router) handleGetSecret(w http.ResponseWriter, r *http.Request) {
	secret, err := router.facade.GetSecret(r.Context(), "test")
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	// TODO: Fix this -> GetSecretResponse
	resp := dto.CreateSecretResponse{
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
