package http

import (
	"errors"
	"net/http"

	"github.com/fuzr0dah/locker/internal/api"

	"github.com/go-chi/render"
)

func respondWithError(w http.ResponseWriter, r *http.Request, err error) {
	var apiErr api.APIError
	if errors.As(err, &apiErr) {
		switch apiErr.Code {
		case api.ErrNotFound:
			render.Status(r, http.StatusNotFound)
		case api.ErrAlreadyExists:
			render.Status(r, http.StatusConflict)
		case api.ErrInvalidInput:
			render.Status(r, http.StatusBadRequest)
		default:
			render.Status(r, http.StatusInternalServerError)
		}
		render.JSON(w, r, apiErr)
		return
	}
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, api.InternalErr)
}

func badRequest(w http.ResponseWriter, r *http.Request, err string) {
	apiErr := api.APIError{Code: api.ErrBadRequest, Message: err}
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, apiErr)
}
