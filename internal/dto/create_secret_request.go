package dto

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateSecretRequest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (csr *CreateSecretRequest) Bind(r *http.Request) error {
	if !json.Valid([]byte(csr.Value)) {
		return fmt.Errorf("value is not valid JSON")
	}
	return nil
}
