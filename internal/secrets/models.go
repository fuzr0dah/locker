package secrets

import (
	"context"
	"time"
)

// Storage defines the interface for secret storage operations
type Storage interface {
	Close() error
	CreateSecret(ctx context.Context, name string, value []byte) (*Secret, error)
	GetSecretById(ctx context.Context, id int64) (*Secret, error)
	UpdateSecret(ctx context.Context, name string, value []byte) (*Secret, error)
	DeleteSecret(ctx context.Context, name string) error
	ListSecrets(ctx context.Context) ([]*Secret, error)
	GetSecretVersions(ctx context.Context, name string, limit int) ([]*SecretVersion, error)
	GetSecretVersion(ctx context.Context, name string, version int) (*SecretVersion, error)
}

// Secret represents the current version of a secret
type Secret struct {
	ID        int64
	Name      string
	Value     []byte
	Version   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// SecretVersion represents a historical version of a secret
type SecretVersion struct {
	ID        int64
	SecretID  int64
	Version   int64
	Value     []byte
	CreatedAt time.Time
}
