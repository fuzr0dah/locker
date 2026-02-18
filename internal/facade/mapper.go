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

func mapToApiSecretVersion(secretVersion *secrets.SecretVersion) *api.SecretVersion {
	if secretVersion == nil {
		return nil
	}
	return &api.SecretVersion{
		ID:        secretVersion.ID,
		SecretID:  secretVersion.SecretID,
		Version:   secretVersion.Version,
		Value:     string(secretVersion.Value),
		CreatedAt: secretVersion.CreatedAt,
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
