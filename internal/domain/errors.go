package domain

import "errors"

var (
	ErrSecretNotFound  = errors.New("secret not found")
	ErrVersionConflict = errors.New("secret versions conflict")
)
