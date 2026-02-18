package secrets

import "github.com/fuzr0dah/locker/internal/db"

func fromDBSecret(s db.Secret) *Secret {
	var valueCopy []byte
	if s.Value != nil {
		valueCopy = make([]byte, len(s.Value))
		copy(valueCopy, s.Value)
	}

	return &Secret{
		ID:        s.ID,
		Name:      s.Name,
		Value:     valueCopy,
		Version:   s.CurrentVersion,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func fromDBSecretVersion(s db.SecretVersion) *SecretVersion {
	var valueCopy []byte
	if s.Value != nil {
		valueCopy = make([]byte, len(s.Value))
		copy(valueCopy, s.Value)
	}

	return &SecretVersion{
		ID:        s.ID,
		SecretID:  s.SecretID,
		Version:   s.Version,
		Value:     valueCopy,
		CreatedAt: s.CreatedAt,
	}
}
