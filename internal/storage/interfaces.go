package storage

import (
	"context"

	"github.com/fuzr0dah/locker/internal/domain"
)

// SecretReader defines read-only operations for secrets
type SecretReader interface {
	GetSecretById(ctx context.Context, id string) (*domain.Secret, error)
	ListSecrets(ctx context.Context) ([]*domain.Secret, error)
	GetSecretVersion(ctx context.Context, id string, version int) (*domain.SecretVersion, error)
	GetSecretVersions(ctx context.Context, id string, limit int) ([]*domain.SecretVersion, error)
}

// SecretWriter defines write operations for secrets
type SecretWriter interface {
	CreateSecret(ctx context.Context, name string, value []byte) (*domain.Secret, error)
	UpdateSecret(ctx context.Context, id, name string, value []byte) (*domain.Secret, error)
	DeleteSecret(ctx context.Context, id string) error
}

// SecretStorage combines read and write operations
type SecretStorage interface {
	SecretReader
	SecretWriter
	Close() error
}

type UnitOfWork interface {
	Begin(ctx context.Context) error
	Commit() error
	Rollback() error
	Secrets() SecretStorage
}
