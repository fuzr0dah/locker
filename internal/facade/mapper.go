package facade

import (
	"errors"

	"github.com/fuzr0dah/locker/internal/api"
	"github.com/fuzr0dah/locker/internal/domain"
)

func mapToApiSecret(secret *domain.Secret) *api.Secret {
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

func mapToApiSecretVersion(secretVersion *domain.SecretVersion) *api.SecretVersion {
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
	case errors.Is(err, domain.ErrSecretNotFound):
		return api.SecretNotFoundErr
	case errors.Is(err, domain.ErrVersionConflict):
		return api.SecretVersionConflictErr
	default:
		return api.InternalErr
	}
}
