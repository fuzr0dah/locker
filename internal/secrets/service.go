package secrets

import (
	"context"
	"fmt"
)

type SecretsService interface {
	Create(ctx context.Context, name string, value string) (*Secret, error)
	GetById(ctx context.Context, id string) (*Secret, error)
	Update(ctx context.Context, id, name, value string) (*Secret, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*Secret, error)
	GetVersion(ctx context.Context, id string, version int) (*SecretVersion, error)
	GetVersions(ctx context.Context, id string, limit int) ([]*SecretVersion, error)
}

// Service handles business logic for secrets
type service struct {
	storage Storage
}

// NewService creates a new secrets service
func NewService(storage Storage) *service {
	return &service{storage: storage}
}

// Create creates a new secret
func (s *service) Create(ctx context.Context, name string, value string) (*Secret, error) {
	secret, err := s.storage.CreateSecret(ctx, name, []byte(value))
	if err != nil {
		return nil, fmt.Errorf("create secret: %w", err)
	}
	return secret, nil
}

// Get retrieves a secret by id
func (s *service) GetById(ctx context.Context, id string) (*Secret, error) {
	secret, err := s.storage.GetSecretById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get secret: %w", err)
	}
	return secret, nil
}

// Update updates a secret value (creates new version)
func (s *service) Update(ctx context.Context, id, name, value string) (*Secret, error) {
	secret, err := s.storage.UpdateSecret(ctx, id, name, []byte(value))
	if err != nil {
		return nil, fmt.Errorf("update secret: %w", err)
	}
	return secret, nil
}

// Delete removes a secret
func (s *service) Delete(ctx context.Context, id string) error {
	if err := s.storage.DeleteSecret(ctx, id); err != nil {
		return fmt.Errorf("delete secret: %w", err)
	}
	return nil
}

// List returns all secrets (without values)
func (s *service) List(ctx context.Context) ([]*Secret, error) {
	secrets, err := s.storage.ListSecrets(ctx)
	if err != nil {
		return nil, fmt.Errorf("list secrets: %w", err)
	}
	return secrets, nil
}

// GetVersion returns a specific version of a secret
func (s *service) GetVersion(ctx context.Context, id string, version int) (*SecretVersion, error) {
	secretVersion, err := s.storage.GetSecretVersion(ctx, id, version)
	if err != nil {
		return nil, fmt.Errorf("get version: %w", err)
	}
	return secretVersion, nil
}

// GetVersions returns version history for a secret
func (s *service) GetVersions(ctx context.Context, id string, limit int) ([]*SecretVersion, error) {
	versions, err := s.storage.GetSecretVersions(ctx, id, limit)
	if err != nil {
		return nil, fmt.Errorf("get versions: %w", err)
	}
	return versions, nil
}
