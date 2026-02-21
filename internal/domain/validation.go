package domain

import "strings"

func ValidateSecretName(name string) error {
	if strings.TrimSpace(name) == "" {
		return ErrNameEmpty
	}
	if len(name) > 255 {
		return ErrNameTooLong
	}
	return nil
}
