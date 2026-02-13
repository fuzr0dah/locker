package facade

import (
	"context"

	"github.com/fuzr0dah/locker/internal/secrets"
)

// SecretsFacade provides high-level operations for secrets management
type SecretsFacade interface {
	CreateSecret(ctx context.Context, name string, value string) (*secrets.Secret, error)
	GetSecretById(ctx context.Context, id int64) (*secrets.Secret, error)
	UpdateSecret(ctx context.Context, name string, value string) (*secrets.Secret, error)
	DeleteSecret(ctx context.Context, name string) error
	ListSecrets(ctx context.Context) ([]*secrets.Secret, error)
	GetSecretVersions(ctx context.Context, name string, limit int) ([]*secrets.SecretVersion, error)
}

// facade implements SecretsFacade
type facade struct {
	service *secrets.Service
}

// NewFacade creates a facade with in-memory storage (for dev mode)
func NewFacade(service *secrets.Service) SecretsFacade {
	return &facade{service: service}
}

func (f *facade) CreateSecret(ctx context.Context, name string, value string) (*secrets.Secret, error) {
	return f.service.Create(ctx, name, value)
}

func (f *facade) GetSecretById(ctx context.Context, id int64) (*secrets.Secret, error) {
	return f.service.GetById(ctx, id)
}

func (f *facade) UpdateSecret(ctx context.Context, name string, value string) (*secrets.Secret, error) {
	return f.service.Update(ctx, name, value)
}

func (f *facade) DeleteSecret(ctx context.Context, name string) error {
	return f.service.Delete(ctx, name)
}

func (f *facade) ListSecrets(ctx context.Context) ([]*secrets.Secret, error) {
	return f.service.List(ctx)
}

func (f *facade) GetSecretVersions(ctx context.Context, name string, limit int) ([]*secrets.SecretVersion, error) {
	return f.service.GetVersions(ctx, name, limit)
}
