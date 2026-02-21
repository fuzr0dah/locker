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

	// 404 Not Found
	if errors.Is(err, domain.ErrSecretNotFound) {
		return api.SecretNotFoundErr
	}

	// 409 Conflict
	if errors.Is(err, domain.ErrVersionConflict) {
		return api.SecretVersionConflictErr
	}
	if errors.Is(err, domain.ErrNameAlreadyExists) {
		return api.APIError{Code: api.ErrAlreadyExists, Message: "secret name already exists"}
	}
	if errors.Is(err, domain.ErrSecretDeleted) {
		return api.APIError{Code: api.ErrConflict, Message: "secret is deleted"}
	}

	// 400 Bad Request
	if errors.Is(err, domain.ErrNameEmpty) || errors.Is(err, domain.ErrNameTooLong) {
		return api.APIError{Code: api.ErrInvalidInput, Message: err.Error()}
	}

	var apiErr api.APIError
	if errors.As(err, &apiErr) {
		return err
	}

	return api.InternalErr
}
