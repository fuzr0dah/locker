package secrets

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/fuzr0dah/locker/internal/domain/crypto"
	"github.com/fuzr0dah/locker/internal/domain/repository"
	"github.com/fuzr0dah/locker/internal/domain/secrets"
	"github.com/fuzr0dah/locker/internal/domain/validation"
)

// secretsService handles business logic for secrets
type secretsService struct {
	envelope   crypto.Envelope
	reader     repository.SecretReader
	uowFactory func() repository.UnitOfWork
	logger     *slog.Logger
}

// NewSecretsService creates a new secrets service
func NewSecretsService(
	envelope crypto.Envelope,
	reader repository.SecretReader,
	uowFactory func() repository.UnitOfWork,
	logger *slog.Logger,
) *secretsService {
	if logger == nil {
		logger = slog.Default()
	}
	return &secretsService{
		envelope:   envelope,
		reader:     reader,
		uowFactory: uowFactory,
		logger:     logger,
	}
}

// Create creates a new secret
func (s *secretsService) Create(ctx context.Context, name string, value string) (*secrets.Secret, error) {
	if err := validation.ValidateSecretName(name); err != nil {
		return nil, err
	}

	cipherValue, err := s.envelope.Seal([]byte(value))
	if err != nil {
		return nil, fmt.Errorf("encrypt value: %w", err)
	}

	uow := s.uowFactory()
	if err := uow.Begin(ctx); err != nil {
		return nil, fmt.Errorf("begin transaction: %w", err)
	}

	var secret *secrets.Secret
	var opErr error

	defer func() {
		if opErr != nil {
			if rbErr := uow.Rollback(); rbErr != nil {
				s.logger.Error("failed to rollback transaction", "error", rbErr)
			}
		}
	}()

	secret, opErr = uow.Writer().CreateSecret(ctx, name, cipherValue)
	if opErr != nil {
		return nil, fmt.Errorf("create secret: %w", opErr)
	}

	if err := uow.Commit(); err != nil {
		return nil, fmt.Errorf("commit transaction: %w", err)
	}

	// Не возвращаем value обратно (клиент его знает), очищаем для безопасности
	secret.Value = nil
	return secret, nil
}

// GetById retrieves a secret by id
func (s *secretsService) GetById(ctx context.Context, id string) (*secrets.Secret, error) {
	secret, err := s.reader.GetSecretById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get secret: %w", err)
	}
	secret.Value, err = s.envelope.Open(secret.Value)
	if err != nil {
		if errors.Is(err, crypto.ErrDecryptionFailed) {
			s.logger.Error("critical: failed to decrypt secret", "secret_id", id, "error", err)
			return nil, fmt.Errorf("secret data integrity compromised: %w", err)
		}
		return nil, err
	}
	return secret, nil
}

// Update updates a secret value (creates new version)
func (s *secretsService) Update(ctx context.Context, id, name, value string) (*secrets.Secret, error) {
	if err := validation.ValidateSecretName(name); err != nil {
		return nil, err
	}

	cipherValue, err := s.envelope.Seal([]byte(value))
	if err != nil {
		return nil, fmt.Errorf("encrypt value: %w", err)
	}

	uow := s.uowFactory()
	if err := uow.Begin(ctx); err != nil {
		return nil, fmt.Errorf("begin transaction: %w", err)
	}

	var secret *secrets.Secret
	var opErr error

	defer func() {
		if opErr != nil {
			if rbErr := uow.Rollback(); rbErr != nil {
				s.logger.Error("failed to rollback transaction", "error", rbErr)
			}
		}
	}()

	secret, opErr = uow.Writer().UpdateSecret(ctx, id, name, cipherValue)
	if opErr != nil {
		return nil, fmt.Errorf("update secret: %w", opErr)
	}

	if opErr = uow.Commit(); opErr != nil {
		return nil, fmt.Errorf("commit transaction: %w", opErr)
	}

	// Не возвращаем value обратно (клиент его знает), очищаем для безопасности
	secret.Value = nil
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
func (s *secretsService) List(ctx context.Context) ([]*secrets.Secret, error) {
	secrets, err := s.reader.ListSecrets(ctx)
	if err != nil {
		return nil, fmt.Errorf("list secrets: %w", err)
	}
	return secrets, nil
}

// GetVersion returns a specific version of a secret
func (s *secretsService) GetVersion(ctx context.Context, id string, version int) (*secrets.SecretVersion, error) {
	secretVersion, err := s.reader.GetSecretVersion(ctx, id, version)
	if err != nil {
		return nil, fmt.Errorf("get version: %w", err)
	}
	return secretVersion, nil
}

// GetVersions returns version history for a secret
func (s *secretsService) GetVersions(ctx context.Context, id string, limit int) ([]*secrets.SecretVersion, error) {
	versions, err := s.reader.GetSecretVersions(ctx, id, limit)
	if err != nil {
		return nil, fmt.Errorf("get versions: %w", err)
	}
	return versions, nil
}
