package server

import (
	"encoding/json"
	"net/http"
)

func decodeRequest(r *http.Request, v interface{ Validate() error }) error {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return v.Validate()
}
