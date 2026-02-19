package sqlite

import (
	"github.com/fuzr0dah/locker/internal/db"
	"github.com/fuzr0dah/locker/internal/domain"
)

func fromDBSecret(s db.Secret) *domain.Secret {
	return &domain.Secret{
		ID:        s.ID,
		Name:      s.Name,
		Version:   s.VersionID.Int64,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func fromGetSecretByIdRow(s db.GetSecretByIdRow) *domain.Secret {
	var valueCopy []byte
	if s.Value != nil {
		valueCopy = make([]byte, len(s.Value))
		copy(valueCopy, s.Value)
	}

	return &domain.Secret{
		ID:        s.ID,
		Name:      s.Name,
		Version:   s.VersionID.Int64,
		Value:     valueCopy,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func fromDBSecretVersion(s db.SecretVersion) *domain.SecretVersion {
	var valueCopy []byte
	if s.Value != nil {
		valueCopy = make([]byte, len(s.Value))
		copy(valueCopy, s.Value)
	}

	return &domain.SecretVersion{
		ID:        s.ID,
		SecretID:  s.SecretID,
		Version:   s.Version,
		Value:     valueCopy,
		CreatedAt: s.CreatedAt,
	}
}
