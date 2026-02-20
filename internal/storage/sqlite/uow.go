package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/fuzr0dah/locker/internal/db"
	"github.com/fuzr0dah/locker/internal/domain"
	"github.com/fuzr0dah/locker/internal/idgen"
	"github.com/fuzr0dah/locker/internal/storage"
)

// unitOfWork implements storage.UnitOfWork interface for SQLite
type unitOfWork struct {
	db      *sql.DB
	tx      *sql.Tx
	queries *db.Queries

	writer storage.SecretWriter
}

// NewUnitOfWork creates a new SQLite unit of work
func NewUnitOfWork(db *sql.DB) storage.UnitOfWork {
	return &unitOfWork{db: db}
}

// Begin starts a new transaction
func (u *unitOfWork) Begin(ctx context.Context) error {
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return err
	}

	u.tx = tx
	u.queries = db.New(tx)

	return nil
}

// Commit commits the transaction
func (u *unitOfWork) Commit() error {
	if u.tx == nil {
		return errors.New("no transaction")
	}
	defer u.cleanup()
	return u.tx.Commit()
}

// Rollback rolls back the transaction
func (u *unitOfWork) Rollback() error {
	if u.tx == nil {
		return nil
	}
	defer u.cleanup()
	return u.tx.Rollback()
}

// Writer returns the secrets writer for this unit of work
func (u *unitOfWork) Writer() storage.SecretWriter {
	if u.queries == nil {
		panic("unitOfWork not started - call Begin() first")
	}
	if u.writer == nil {
		u.writer = &txStorage{secretReader: &secretReader{queries: u.queries}}
	}
	return u.writer
}

func (u *unitOfWork) cleanup() {
	u.tx = nil
	u.queries = nil
	u.writer = nil
}

// txStorage is a storage implementation that operates within a transaction
type txStorage struct {
	*secretReader
}

func (s *txStorage) Close() error {
	return nil
}

func (s *txStorage) CreateSecret(ctx context.Context, name string, value []byte) (*domain.Secret, error) {
	secretID, err := idgen.SecretID()
	if err != nil {
		return nil, fmt.Errorf("generate secret id: %w", err)
	}

	_, err = s.secretReader.queries.CreateSecret(ctx, db.CreateSecretParams{ID: secretID, Name: name})
	if err != nil {
		return nil, fmt.Errorf("create secret record: %w", err)
	}

	secretVersion, err := s.secretReader.queries.CreateInitialVersion(ctx, db.CreateInitialVersionParams{
		SecretID: secretID,
		Value:    value,
	})
	if err != nil {
		return nil, fmt.Errorf("create initial version: %w", err)
	}

	secret, err := s.secretReader.queries.InsertVersionIntoSecret(ctx, db.InsertVersionIntoSecretParams{
		ID:        secretID,
		VersionID: sql.NullInt64{Int64: secretVersion.ID, Valid: true},
	})
	if err != nil {
		return nil, fmt.Errorf("link version to secret: %w", err)
	}

	return fromDBSecret(secret), nil
}

func (s *txStorage) UpdateSecret(ctx context.Context, id, name string, value []byte) (*domain.Secret, error) {
	currentVersion, err := s.secretReader.queries.GetLastVersionForSecretId(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrSecretNotFound
		}
		return nil, fmt.Errorf("get current version: %w", err)
	}

	newVersion, err := s.secretReader.queries.CreateNextVersion(ctx, db.CreateNextVersionParams{
		SecretID: id,
		Version:  currentVersion.Version + 1,
		Value:    value,
	})
	if err != nil {
		return nil, fmt.Errorf("create new version: %w", err)
	}

	updated, err := s.secretReader.queries.UpdateSecret(ctx, db.UpdateSecretParams{
		ID:           id,
		Name:         name,
		VersionID:    sql.NullInt64{Int64: newVersion.ID, Valid: true},
		OldVersionID: sql.NullInt64{Int64: currentVersion.ID, Valid: true},
	})
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrVersionConflict
	}
	if err != nil {
		return nil, fmt.Errorf("update secret: %w", err)
	}

	return fromDBSecret(updated), nil
}

func (s *txStorage) DeleteSecret(ctx context.Context, id string) error {
	err := s.secretReader.queries.DeleteSecret(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.ErrSecretNotFound
	}
	if err != nil {
		return fmt.Errorf("delete secret: %w", err)
	}
	return nil
}
