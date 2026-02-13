package secrets

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/fuzr0dah/locker/internal/db/migrations"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

// InMemoryStorage implements Storage interface using SQLite in-memory mode
type InMemoryStorage struct {
	db *sql.DB
}

// NewInMemoryStorage creates and initializes a new in-memory storage instance
func NewInMemoryStorage() (*InMemoryStorage, error) {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared&_fk=on")
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	if err := migrations.RunMigrations(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("run migrations: %w", err)
	}

	return &InMemoryStorage{db: db}, nil
}

// Close closes the storage connection
func (s *InMemoryStorage) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

// CreateSecret creates a new secret with the given name and value
func (s *InMemoryStorage) CreateSecret(ctx context.Context, name string, value []byte) (*Secret, error) {
	// TODO: implement using sqlc queries
	return nil, errors.New("not implemented")
}

// GetSecret retrieves a secret by name
func (s *InMemoryStorage) GetSecret(ctx context.Context, name string) (*Secret, error) {
	// TODO: implement using sqlc queries
	return nil, errors.New("not implemented")
}

// UpdateSecret updates the value of an existing secret
func (s *InMemoryStorage) UpdateSecret(ctx context.Context, name string, value []byte) (*Secret, error) {
	// TODO: implement using sqlc queries
	return nil, errors.New("not implemented")
}

// DeleteSecret removes a secret and all its versions (CASCADE)
func (s *InMemoryStorage) DeleteSecret(ctx context.Context, name string) error {
	// TODO: implement using sqlc queries
	return errors.New("not implemented")
}

// ListSecrets returns all secrets (without values for performance)
func (s *InMemoryStorage) ListSecrets(ctx context.Context) ([]*Secret, error) {
	// TODO: implement using sqlc queries
	return nil, errors.New("not implemented")
}

// GetSecretVersions returns version history for a secret
func (s *InMemoryStorage) GetSecretVersions(ctx context.Context, name string, limit int) ([]*SecretVersion, error) {
	// TODO: implement using sqlc queries
	return nil, errors.New("not implemented")
}

// GetSecretVersion returns a specific version of a secret
func (s *InMemoryStorage) GetSecretVersion(ctx context.Context, name string, version int) (*SecretVersion, error) {
	// TODO: implement using sqlc queries
	return nil, errors.New("not implemented")
}
