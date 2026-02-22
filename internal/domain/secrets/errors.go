package secrets

import "errors"

var (
	ErrSecretNotFound = errors.New("secret not found")

	ErrVersionConflict   = errors.New("secret versions conflict")
	ErrNameAlreadyExists = errors.New("secret name already exists")
	ErrSecretDeleted     = errors.New("secret is deleted")
)
