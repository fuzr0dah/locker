package http

import (
	"encoding/json"
	"net/http"

	"github.com/fuzr0dah/locker/internal/api"
)

type Validator interface {
	Validate() error
}

func decodeRequest(r *http.Request, v any) error {
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return api.APIError{
			Code:    api.ErrInvalidInput,
			Message: "invalid JSON: " + err.Error(),
		}
	}

	if v, ok := v.(Validator); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}
