package facade

import (
	"context"

	"github.com/fuzr0dah/locker/internal/secrets"
)

// SecretsFacade provides high-level operations for secrets management
type SecretsFacade interface {
	CreateSecret(ctx context.Context, name string, value []byte) (*secrets.Secret, error)
	GetSecret(ctx context.Context, name string) (*secrets.Secret, error)
	UpdateSecret(ctx context.Context, name string, value []byte) (*secrets.Secret, error)
	DeleteSecret(ctx context.Context, name string) error
	ListSecrets(ctx context.Context) ([]*secrets.Secret, error)
	GetSecretVersions(ctx context.Context, name string, limit int) ([]*secrets.SecretVersion, error)
	Close() error
}

// facade implements SecretsFacade
type facade struct {
	service *secrets.Service
}

// NewInMemoryFacade creates a facade with in-memory storage (for dev mode)
func NewInMemoryFacade() (SecretsFacade, error) {
	storage, err := secrets.NewInMemoryStorage()
	if err != nil {
		return nil, err
	}

	service := secrets.NewService(storage)
	return &facade{service: service}, nil
}

func (f *facade) CreateSecret(ctx context.Context, name string, value []byte) (*secrets.Secret, error) {
	return f.service.Create(ctx, name, value)
}

func (f *facade) GetSecret(ctx context.Context, name string) (*secrets.Secret, error) {
	return f.service.Get(ctx, name)
}

func (f *facade) UpdateSecret(ctx context.Context, name string, value []byte) (*secrets.Secret, error) {
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

func (f *facade) Close() error {
	return f.service.Close()
}
