package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/fuzr0dah/locker/internal/db"
	"github.com/fuzr0dah/locker/internal/domain"
)

// secretReader implements read-only operations for secrets
type secretReader struct {
	queries *db.Queries
}

// NewSecretReader creates a new read-only storage instance
func NewSecretReader(conn *sql.DB) *secretReader {
	return &secretReader{
		queries: db.New(conn),
	}
}

// GetSecretById retrieves a secret by id
func (r *secretReader) GetSecretById(ctx context.Context, id string) (*domain.Secret, error) {
	secret, err := r.queries.GetSecretById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrSecretNotFound
		}
		return nil, fmt.Errorf("get secret: %w", err)
	}
	return fromGetSecretByIdRow(secret), nil
}

// ListSecrets returns all secrets
func (r *secretReader) ListSecrets(ctx context.Context) ([]*domain.Secret, error) {
	list, err := r.queries.ListSecrets(ctx)
	if err != nil {
		return nil, fmt.Errorf("list secrets: %w", err)
	}

	secretList := make([]*domain.Secret, len(list))
	for i := range list {
		secretList[i] = fromDBSecret(list[i])
	}

	return secretList, nil
}

// GetSecretVersion returns a specific version of a secret
func (r *secretReader) GetSecretVersion(ctx context.Context, id string, version int) (*domain.SecretVersion, error) {
	secretVersion, err := r.queries.GetSecretVersion(ctx, db.GetSecretVersionParams{
		SecretID: id,
		Version:  int64(version),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrSecretNotFound
		}
		return nil, fmt.Errorf("get secret version: %w", err)
	}
	return fromDBSecretVersion(secretVersion), nil
}

// GetSecretVersions returns version history for a secret
func (r *secretReader) GetSecretVersions(ctx context.Context, id string, limit int) ([]*domain.SecretVersion, error) {
	// TODO: implement using sqlc queries
	return nil, errors.New("not implemented")
}
