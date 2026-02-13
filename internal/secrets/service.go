package secrets

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrSecretNotFound = errors.New("secret not found")
	ErrSecretExists   = errors.New("secret already exists")
)

// Service handles business logic for secrets
type Service struct {
	storage Storage
}

// NewService creates a new secrets service
func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

// Create creates a new secret
func (s *Service) Create(ctx context.Context, name string, value string) (*Secret, error) {
	secret, err := s.storage.CreateSecret(ctx, name, []byte(value))
	if err != nil {
		return nil, fmt.Errorf("create secret: %w", err)
	}
	return secret, nil
}

// Get retrieves a secret by id
func (s *Service) GetById(ctx context.Context, id int64) (*Secret, error) {
	secret, err := s.storage.GetSecretById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get secret: %w", err)
	}
	return secret, nil
}

// Update updates a secret value (creates new version)
func (s *Service) Update(ctx context.Context, name string, value string) (*Secret, error) {
	secret, err := s.storage.UpdateSecret(ctx, name, []byte(value))
	if err != nil {
		return nil, fmt.Errorf("update secret: %w", err)
	}
	return secret, nil
}

// Delete removes a secret
func (s *Service) Delete(ctx context.Context, name string) error {
	if err := s.storage.DeleteSecret(ctx, name); err != nil {
		return fmt.Errorf("delete secret: %w", err)
	}
	return nil
}

// List returns all secrets (without values)
func (s *Service) List(ctx context.Context) ([]*Secret, error) {
	secrets, err := s.storage.ListSecrets(ctx)
	if err != nil {
		return nil, fmt.Errorf("list secrets: %w", err)
	}
	return secrets, nil
}

// GetVersions returns version history for a secret
func (s *Service) GetVersions(ctx context.Context, name string, limit int) ([]*SecretVersion, error) {
	versions, err := s.storage.GetSecretVersions(ctx, name, limit)
	if err != nil {
		return nil, fmt.Errorf("get versions: %w", err)
	}
	return versions, nil
}

// GetVersion returns a specific version of a secret
func (s *Service) GetVersion(ctx context.Context, name string, version int) (*SecretVersion, error) {
	v, err := s.storage.GetSecretVersion(ctx, name, version)
	if err != nil {
		return nil, fmt.Errorf("get version: %w", err)
	}
	return v, nil
}
