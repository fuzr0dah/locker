package repository

import (
	"context"

	"github.com/fuzr0dah/locker/internal/domain/secrets"
)

type SecretWriter interface {
	CreateSecret(ctx context.Context, name string, value []byte) (*secrets.Secret, error)
	UpdateSecret(ctx context.Context, id, name string, value []byte) (*secrets.Secret, error)
	DeleteSecret(ctx context.Context, id string) error
}
