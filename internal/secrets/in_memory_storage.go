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

// inMemoryStorage implements Storage interface using SQLite in-memory mode
type inMemoryStorage struct {
	db      *sql.DB
	queries *db.Queries
}

// OpenDB opens a new SQLite in-memory database connection
func OpenDB() (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", "file::memory:?cache=shared&_fk=on&_journal_mode=WAL&_busy_timeout=5000")

	conn.SetMaxOpenConns(1)
	conn.SetMaxIdleConns(1)

	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}
	return conn, nil
}

// NewStorage creates a new storage instance from an existing database connection
func NewStorage(conn *sql.DB) *inMemoryStorage {
	return &inMemoryStorage{
		db:      conn,
		queries: db.New(conn),
	}
}

// Close closes the storage connection
func (s *inMemoryStorage) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

// CreateSecret creates a new secret with the given name and value
func (s *inMemoryStorage) CreateSecret(ctx context.Context, name string, value []byte) (*Secret, error) {
	// TODO retry if generateSecretID make id with collisions
	newId, err := generateSecretID()
	if err != nil {
		return nil, fmt.Errorf("create secret: %w", err)
	}
	params := db.CreateSecretParams{ID: newId, Name: name, Value: value}
	secret, err := s.queries.CreateSecret(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("create secret: %w", err)
	}
	return fromDBSecret(secret), nil
}

// GetSecretById retrieves a secret by id
func (s *inMemoryStorage) GetSecretById(ctx context.Context, id string) (*Secret, error) {
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
func (s *inMemoryStorage) UpdateSecret(ctx context.Context, id, name string, value []byte) (*Secret, error) {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	current, err := s.queries.WithTx(tx).GetSecretById(ctx, id)
	if err != nil {
		return nil, err
	}

	_, err = s.queries.WithTx(tx).CreateSecretVersion(ctx, db.CreateSecretVersionParams{
		SecretID: current.ID,
		Version:  current.CurrentVersion,
		Value:    current.Value,
	})
	if err != nil {
		return nil, err
	}

	updated, err := s.queries.WithTx(tx).UpdateSecret(ctx, db.UpdateSecretParams{
		ID:             id,
		Name:           name,
		Value:          value,
		CurrentVersion: current.CurrentVersion,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrVersionConflict
	}
	if err != nil {
		return nil, fmt.Errorf("update: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return fromDBSecret(updated), nil
}

// DeleteSecret removes a secret and all its versions (CASCADE)
func (s *inMemoryStorage) DeleteSecret(ctx context.Context, id string) error {
	return s.queries.DeleteSecret(ctx, id)
}

// ListSecrets returns all secrets
func (s *inMemoryStorage) ListSecrets(ctx context.Context) ([]*Secret, error) {
	list, err := s.queries.ListSecrets(ctx)
	if err != nil {
		return nil, fmt.Errorf("list secrets: %w", err)
	}

	secrets := make([]*Secret, len(list))
	for i := range list {
		secrets[i] = fromDBSecret(list[i])
	}

	return secrets, nil
}

// GetSecretVersions returns version history for a secret
func (s *inMemoryStorage) GetSecretVersion(ctx context.Context, id string, version int) (*SecretVersion, error) {
	secretVersion, err := s.queries.GetSecretVersion(ctx, db.GetSecretVersionParams{SecretID: id, Version: int64(version)})
	if err != nil {
		return nil, fmt.Errorf("get version: %w", err)
	}
	return fromDBSecretVersion(secretVersion), nil
}

// GetSecretVersion returns a specific version of a secret
func (s *inMemoryStorage) GetSecretVersions(ctx context.Context, id string, limit int) ([]*SecretVersion, error) {
	// TODO: implement using sqlc queries
	return nil, errors.New("not implemented")
}
