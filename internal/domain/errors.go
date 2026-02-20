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
	ErrInvalidName   = errors.New("invalid secret name")
	ErrInvalidValue  = errors.New("invalid secret value")
	ErrNameTooLong   = errors.New("secret name too long")
	ErrValueTooLarge = errors.New("secret value too large")
	ErrNameEmpty     = errors.New("secret name cannot be empty")
)
