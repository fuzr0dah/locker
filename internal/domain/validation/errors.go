package validation

import "errors"

var (
	ErrNameTooLong = errors.New("secret name too long")
	ErrNameEmpty   = errors.New("secret name cannot be empty")
)
