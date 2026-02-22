package secrets

import "context"

type SecretsService interface {
	Create(ctx context.Context, name string, value string) (*Secret, error)
	GetById(ctx context.Context, id string) (*Secret, error)
	Update(ctx context.Context, id, name, value string) (*Secret, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*Secret, error)
	GetVersion(ctx context.Context, id string, version int) (*SecretVersion, error)
	GetVersions(ctx context.Context, id string, limit int) ([]*SecretVersion, error)
}
