package facade

import (
	"context"
	"strings"

	"github.com/fuzr0dah/locker/internal/api"
	"github.com/fuzr0dah/locker/internal/secrets"
)

// SecretsFacade provides high-level operations for secrets management
type SecretsFacade interface {
	CreateSecret(ctx context.Context, name string, value string) (*api.Secret, error)
	GetSecretById(ctx context.Context, id string) (*api.Secret, error)
	UpdateSecret(ctx context.Context, id, name, value string) (*api.Secret, error)
	DeleteSecret(ctx context.Context, name string) error
	ListSecrets(ctx context.Context) ([]*api.Secret, error)
	GetSecretVersions(ctx context.Context, name string, limit int) ([]*secrets.SecretVersion, error)
}

// facade implements SecretsFacade
type facade struct {
	service secrets.SecretsService
}

// NewFacade creates a facade with in-memory storage (for dev mode)
func NewFacade(service secrets.SecretsService) SecretsFacade {
	return &facade{service: service}
}

func (f *facade) CreateSecret(ctx context.Context, name string, value string) (*api.Secret, error) {
	secret, err := f.service.Create(ctx, name, value)
	if err != nil {
		return nil, mapToApiError(err)
	}
	return mapToApiSecret(secret), nil
}

func (f *facade) GetSecretById(ctx context.Context, id string) (*api.Secret, error) {
	secret, err := f.service.GetById(ctx, id)
	if err != nil {
		return nil, mapToApiError(err)
	}
	return mapToApiSecret(secret), nil
}

func (f *facade) UpdateSecret(ctx context.Context, id, name, value string) (*api.Secret, error) {
	secret, err := f.service.Update(ctx, id, name, value)
	if err != nil {
		// TODO fix this swap to domain error
		if strings.Contains(err.Error(), "version conflict") {
			return nil, api.APIError{
				Code:    api.ErrConflict,
				Message: "secret was modified by another request, retry",
			}
		}
		return nil, mapToApiError(err)
	}
	return mapToApiSecret(secret), nil
}

func (f *facade) DeleteSecret(ctx context.Context, name string) error {
	return mapToApiError(f.service.Delete(ctx, name))
}

func (f *facade) ListSecrets(ctx context.Context) ([]*api.Secret, error) {
	list, err := f.service.List(ctx)
	if err != nil {
		return nil, mapToApiError(err)
	}
	mapped := make([]*api.Secret, len(list))
	for i := range list {
		mapped[i] = mapToApiSecret(list[i])
	}
	return mapped, nil
}

func (f *facade) GetSecretVersions(ctx context.Context, name string, limit int) ([]*secrets.SecretVersion, error) {
	return f.service.GetVersions(ctx, name, limit)
}
