package secrets

import (
	"context"
	"errors"
	"time"
)

// Storage defines the interface for secret storage operations
type Storage interface {
	Close() error
	CreateSecret(ctx context.Context, name string, value []byte) (*Secret, error)
	GetSecretById(ctx context.Context, id string) (*Secret, error)
	UpdateSecret(ctx context.Context, id, name string, value []byte) (*Secret, error)
	DeleteSecret(ctx context.Context, id string) error
	ListSecrets(ctx context.Context) ([]*Secret, error)
	GetSecretVersion(ctx context.Context, id string, version int) (*SecretVersion, error)
	GetSecretVersions(ctx context.Context, id string, limit int) ([]*SecretVersion, error)
}

// Secret represents the current version of a secret
type Secret struct {
	ID        string
	Name      string
	Value     []byte
	Version   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// SecretVersion represents a historical version of a secret
type SecretVersion struct {
	ID        int64
	SecretID  string
	Version   int64
	Value     []byte
	CreatedAt time.Time
}

var (
	ErrSecretNotFound  = errors.New("secret not found")
	ErrVersionConflict = errors.New("secret versions conflict")
)
