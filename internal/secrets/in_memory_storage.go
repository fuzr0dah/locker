package secrets

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/fuzr0dah/locker/internal/db"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

// InMemoryStorage implements Storage interface using SQLite in-memory mode
type InMemoryStorage struct {
	db      *sql.DB
	queries *db.Queries
}

// OpenDB opens a new SQLite in-memory database connection
func OpenDB() (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", "file::memory:?cache=shared&_fk=on")
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}
	return conn, nil
}

// NewStorage creates a new storage instance from an existing database connection
func NewStorage(conn *sql.DB) *InMemoryStorage {
	return &InMemoryStorage{
		db:      conn,
		queries: db.New(conn),
	}
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
	params := db.CreateSecretParams{Name: name, Value: value}
	secret, err := s.queries.CreateSecret(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("create secret: %w", err)
	}
	return fromDBSecret(secret), nil
}

// GetSecretById retrieves a secret by id
func (s *InMemoryStorage) GetSecretById(ctx context.Context, id int64) (*Secret, error) {
	secret, err := s.queries.GetSecretById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrSecretNotFound
		}
		return nil, fmt.Errorf("get secret: %w", err)
	}
	return fromDBSecret(secret), nil
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
	dto, err := s.queries.ListSecrets(ctx)
	if err != nil {
		return nil, fmt.Errorf("list secrets: %w", err)
	}
	secrets := make([]*Secret, len(dto))

	for i := range dto {
		secrets[i] = fromDBSecret(dto[i])
	}
	return secrets, nil
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
