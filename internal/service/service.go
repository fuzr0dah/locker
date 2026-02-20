package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/fuzr0dah/locker/internal/domain"
	"github.com/fuzr0dah/locker/internal/storage"
)

type SecretsService interface {
	Create(ctx context.Context, name string, value string) (*domain.Secret, error)
	GetById(ctx context.Context, id string) (*domain.Secret, error)
	Update(ctx context.Context, id, name, value string) (*domain.Secret, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*domain.Secret, error)
	GetVersion(ctx context.Context, id string, version int) (*domain.SecretVersion, error)
	GetVersions(ctx context.Context, id string, limit int) ([]*domain.SecretVersion, error)
}

// secretsService handles business logic for secrets
type secretsService struct {
	reader     storage.SecretReader
	uowFactory func() storage.UnitOfWork
	logger     *slog.Logger
}

// NewSecretsService creates a new secrets service
func NewSecretsService(reader storage.SecretReader, uowFactory func() storage.UnitOfWork, logger *slog.Logger) SecretsService {
	if logger == nil {
		logger = slog.Default()
	}
	return &secretsService{
		reader:     reader,
		uowFactory: uowFactory,
		logger:     logger,
	}
}

// Create creates a new secret
func (s *secretsService) Create(ctx context.Context, name string, value string) (*domain.Secret, error) {
	uow := s.uowFactory()
	if err := uow.Begin(ctx); err != nil {
		return nil, fmt.Errorf("begin transaction: %w", err)
	}

	var secret *domain.Secret
	var opErr error

	defer func() {
		if opErr != nil {
			if rbErr := uow.Rollback(); rbErr != nil {
				s.logger.Error("failed to rollback transaction", "error", rbErr)
			}
		}
	}()

	secret, opErr = uow.Writer().CreateSecret(ctx, name, []byte(value))
	if opErr != nil {
		return nil, fmt.Errorf("create secret: %w", opErr)
	}

	if err := uow.Commit(); err != nil {
		return nil, fmt.Errorf("commit transaction: %w", err)
	}

	return secret, nil
}

// GetById retrieves a secret by id
func (s *secretsService) GetById(ctx context.Context, id string) (*domain.Secret, error) {
	secret, err := s.reader.GetSecretById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get secret: %w", err)
	}
	return secret, nil
}

// Update updates a secret value (creates new version)
func (s *secretsService) Update(ctx context.Context, id, name, value string) (*domain.Secret, error) {
	uow := s.uowFactory()
	if err := uow.Begin(ctx); err != nil {
		return nil, fmt.Errorf("begin transaction: %w", err)
	}

	var secret *domain.Secret
	var opErr error

	defer func() {
		if opErr != nil {
			if rbErr := uow.Rollback(); rbErr != nil {
				s.logger.Error("failed to rollback transaction", "error", rbErr)
			}
		}
	}()

	secret, opErr = uow.Writer().UpdateSecret(ctx, id, name, []byte(value))
	if opErr != nil {
		return nil, fmt.Errorf("update secret: %w", opErr)
	}

	if opErr = uow.Commit(); opErr != nil {
		return nil, fmt.Errorf("commit transaction: %w", opErr)
	}

	return secret, nil
}

// Delete removes a secret
func (s *secretsService) Delete(ctx context.Context, id string) error {
	uow := s.uowFactory()
	if err := uow.Begin(ctx); err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	var opErr error

	defer func() {
		if opErr != nil {
			if rbErr := uow.Rollback(); rbErr != nil {
				s.logger.Error("failed to rollback transaction", "error", rbErr)
			}
		}
	}()

	opErr = uow.Writer().DeleteSecret(ctx, id)
	if opErr != nil {
		return fmt.Errorf("delete secret: %w", opErr)
	}

	if err := uow.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

// List returns all secrets (without values)
func (s *secretsService) List(ctx context.Context) ([]*domain.Secret, error) {
	secrets, err := s.reader.ListSecrets(ctx)
	if err != nil {
		return nil, fmt.Errorf("list secrets: %w", err)
	}
	return secrets, nil
}

// GetVersion returns a specific version of a secret
func (s *secretsService) GetVersion(ctx context.Context, id string, version int) (*domain.SecretVersion, error) {
	secretVersion, err := s.reader.GetSecretVersion(ctx, id, version)
	if err != nil {
		return nil, fmt.Errorf("get version: %w", err)
	}
	return secretVersion, nil
}

// GetVersions returns version history for a secret
func (s *secretsService) GetVersions(ctx context.Context, id string, limit int) ([]*domain.SecretVersion, error) {
	versions, err := s.reader.GetSecretVersions(ctx, id, limit)
	if err != nil {
		return nil, fmt.Errorf("get versions: %w", err)
	}
	return versions, nil
}
