package domain

import "errors"

var (
	// Not found errors (404)
	ErrSecretNotFound = errors.New("secret not found")

	// Conflict errors (409)
	ErrVersionConflict   = errors.New("secret versions conflict")
	ErrNameAlreadyExists = errors.New("secret name already exists")
	ErrSecretDeleted     = errors.New("secret is deleted")

	// Validation errors (400)
	ErrNameTooLong = errors.New("secret name too long")
	ErrNameEmpty   = errors.New("secret name cannot be empty")
)
