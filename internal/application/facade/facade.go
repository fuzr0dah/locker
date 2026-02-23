package facade

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/fuzr0dah/locker/internal/api"
	"github.com/fuzr0dah/locker/internal/domain/secrets"
)

// SecretsFacade provides high-level operations for secrets management
type SecretsFacade interface {
	CreateSecret(ctx context.Context, name string, value string) (*api.Secret, error)
	GetSecretById(ctx context.Context, id string) (*api.Secret, error)
	UpdateSecret(ctx context.Context, id, name, value string) (*api.Secret, error)
	DeleteSecret(ctx context.Context, id string) error
	ListSecrets(ctx context.Context) ([]*api.Secret, error)
	GetSecretVersion(ctx context.Context, id string, version int) (*api.SecretVersion, error)
	GetSecretVersions(ctx context.Context, id string, limit int) ([]*api.SecretVersion, error)
}

// facade implements SecretsFacade
type facade struct {
	service     secrets.SecretsService
	logger      *slog.Logger
	auditLogger *slog.Logger
}

// NewFacade creates a facade
func NewFacade(service secrets.SecretsService, logger, auditLogger *slog.Logger) *facade {
	return &facade{
		service:     service,
		logger:      logger,
		auditLogger: auditLogger,
	}
}

func (f *facade) CreateSecret(ctx context.Context, name string, value string) (*api.Secret, error) {
	logger := f.logger.With("operation", "create_secret")
	logger.Info("creating secret", "name", name)
	secret, err := f.service.Create(ctx, name, value)
	if err != nil {
		logger.Error("failed to create secret", "error", err)
		return nil, mapToApiError(err)
	}
	logger.Info("secret created", "id", secret.ID)
	return mapToApiSecret(secret), nil
}

func (f *facade) GetSecretById(ctx context.Context, id string) (*api.Secret, error) {
	logger := f.logger.With("operation", "get_secret")
	logger.Info("getting secret", "id", id)
	secret, err := f.service.GetById(ctx, id)
	if err != nil {
		logger.Error("failed to get secret", "error", err)
		return nil, mapToApiError(err)
	}
	logger.Info("secret retrieved", "id", id, "version", secret.Version)
	return mapToApiSecret(secret), nil
}

func (f *facade) UpdateSecret(ctx context.Context, id, name, value string) (*api.Secret, error) {
	logger := f.logger.With("operation", "update_secret")
	logger.Info("updating secret", "id", id, "name", name)
	secret, err := f.service.Update(ctx, id, name, value)
	if err != nil {
		logger.Error("failed to update secret", "error", err)
		return nil, mapToApiError(err)
	}
	logger.Info("secret updated", "id", secret.ID, "version", secret.Version)
	return mapToApiSecret(secret), nil
}

func (f *facade) DeleteSecret(ctx context.Context, id string) error {
	logger := f.logger.With("operation", "delete_secret")
	logger.Info("deleting secret", "id", id)
	err := f.service.Delete(ctx, id)
	if err != nil {
		logger.Error("failed to delete secret", "error", err)
		return mapToApiError(err)
	}
	logger.Info("secret deleted", "id", id)
	return nil
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

func (f *facade) GetSecretVersion(ctx context.Context, id string, version int) (*api.SecretVersion, error) {
	logger := f.logger.With("operation", "get_secret_version")
	logger.Info("retrieving secret version", "id", id, "version", version)
	secretVersion, err := f.service.GetVersion(ctx, id, version)
	if err != nil {
		logger.Error("failed to get secret version", "error", err)
		return nil, mapToApiError(err)
	}
	logger.Info("secret version retrieved", "id", secretVersion.ID, "secret_id", id, "version", version)
	return mapToApiSecretVersion(secretVersion), nil
}

func (f *facade) GetSecretVersions(ctx context.Context, id string, limit int) ([]*api.SecretVersion, error) {
	return nil, fmt.Errorf("not implemented")
}
