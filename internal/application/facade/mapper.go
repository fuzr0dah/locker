package facade

import (
	"errors"

	"github.com/fuzr0dah/locker/internal/api"
	"github.com/fuzr0dah/locker/internal/domain/crypto"
	"github.com/fuzr0dah/locker/internal/domain/secrets"
	"github.com/fuzr0dah/locker/internal/domain/validation"
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

	// 404 Not Found
	if errors.Is(err, secrets.ErrSecretNotFound) {
		return api.SecretNotFoundErr
	}

	// 409 Conflict
	if errors.Is(err, secrets.ErrVersionConflict) {
		return api.SecretVersionConflictErr
	}
	if errors.Is(err, secrets.ErrNameAlreadyExists) {
		return api.SecretNameAlreadyExistsErr
	}
	if errors.Is(err, secrets.ErrSecretDeleted) {
		return api.SecretDeletedErr
	}

	// 400 Bad Request
	if errors.Is(err, validation.ErrNameEmpty) || errors.Is(err, validation.ErrNameTooLong) {
		return api.APIError{Code: api.ErrInvalidInput, Message: err.Error()}
	}
	if errors.Is(err, crypto.ErrInvalidCiphertext) {
		return api.InvalidCiphertextErr
	}

	// 500 Internal Server Error
	if errors.Is(err, crypto.ErrDecryptionFailed) {
		return api.InternalErr
	}

	var apiErr api.APIError
	if errors.As(err, &apiErr) {
		return err
	}

	return api.InternalErr
}
