package sqlite

import (
	"github.com/fuzr0dah/locker/internal/domain/secrets"
	"github.com/fuzr0dah/locker/internal/infrastructure/storage/sqlite/sqlitegen"
)

func fromSQLiteSecret(s sqlitegen.Secret) *secrets.Secret {
	return &secrets.Secret{
		ID:        s.ID,
		Name:      s.Name,
		Version:   s.VersionID.Int64,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func fromSQLiteGetSecretByIdRow(s sqlitegen.GetSecretByIdRow) *secrets.Secret {
	var valueCopy []byte
	if s.Value != nil {
		valueCopy = make([]byte, len(s.Value))
		copy(valueCopy, s.Value)
	}

	return &secrets.Secret{
		ID:        s.ID,
		Name:      s.Name,
		Version:   s.VersionID.Int64,
		Value:     valueCopy,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func fromSQLiteSecretVersion(s sqlitegen.SecretVersion) *secrets.SecretVersion {
	var valueCopy []byte
	if s.Value != nil {
		valueCopy = make([]byte, len(s.Value))
		copy(valueCopy, s.Value)
	}

	return &secrets.SecretVersion{
		ID:        s.ID,
		SecretID:  s.SecretID,
		Version:   s.Version,
		Value:     valueCopy,
		CreatedAt: s.CreatedAt,
	}
}
