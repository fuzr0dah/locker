package facade

import (
	"errors"

	"github.com/fuzr0dah/locker/internal/api"
	"github.com/fuzr0dah/locker/internal/secrets"
)

func mapToApiSecret(secret *secrets.Secret) *api.Secret {
	if secret == nil {
		return nil
	}
	return &api.Secret{
		ID:        secret.ID,
		Name:      secret.Name,
		Value:     string(secret.Value),
		Version:   secret.Version,
		CreatedAt: secret.CreatedAt,
		UpdatedAt: secret.UpdatedAt,
	}
}

func mapToApiError(err error) error {
	if err == nil {
		return nil
	}
	switch {
	case errors.Is(err, secrets.ErrSecretNotFound):
		return api.SecretNotFoundErr
	case errors.Is(err, secrets.ErrVersionConflict):
		return api.SecretVersionConflictErr
	default:
		return api.InternalErr
	}
}
