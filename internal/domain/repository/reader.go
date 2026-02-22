package repository

import (
	"context"

	"github.com/fuzr0dah/locker/internal/domain/secrets"
)

type SecretReader interface {
	GetSecretById(ctx context.Context, id string) (*secrets.Secret, error)
	ListSecrets(ctx context.Context) ([]*secrets.Secret, error)
	GetSecretVersion(ctx context.Context, id string, version int) (*secrets.SecretVersion, error)
	GetSecretVersions(ctx context.Context, id string, limit int) ([]*secrets.SecretVersion, error)
}
