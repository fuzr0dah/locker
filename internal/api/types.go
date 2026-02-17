package api

import (
	"encoding/json"
	"time"
)

type Secret struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Value     string    `json:"value,omitempty"`
	Version   int64     `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateSecretRequest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (r *CreateSecretRequest) Validate() error {
	if r.Name == "" {
		return APIError{Code: ErrBadRequest, Message: "name is required"}
	}
	if !json.Valid([]byte(r.Value)) {
		return APIError{Code: ErrBadRequest, Message: "must be valid JSON"}
	}
	return nil
}

type UpdateSecretRequest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// TODO add validation function for name and value
func (r *UpdateSecretRequest) Validate() error {
	if r.Name == "" {
		// TODO add function for ErrBadRequest with return APIError
		return APIError{Code: ErrBadRequest, Message: "name is required"}
	}
	if !json.Valid([]byte(r.Value)) {
		return APIError{Code: ErrBadRequest, Message: "must be valid JSON"}
	}
	return nil
}
