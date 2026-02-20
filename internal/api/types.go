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

type SecretVersion struct {
	ID        int64     `json:"id"`
	SecretID  string    `json:"secret_id"`
	Version   int64     `json:"version"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateSecretRequest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (r *CreateSecretRequest) Validate() error {
	if r.Value != "" && !json.Valid([]byte(r.Value)) {
		return APIError{Code: ErrInvalidInput, Message: "value must be valid JSON"}
	}
	return nil
}

type UpdateSecretRequest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (r *UpdateSecretRequest) Validate() error {
	if r.Value != "" && !json.Valid([]byte(r.Value)) {
		return APIError{Code: ErrInvalidInput, Message: "value must be valid JSON"}
	}
	return nil
}
