package domain

import "time"

// Secret represents the current version of a secret
type Secret struct {
	ID        string
	Name      string
	Value     []byte
	Version   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// SecretVersion represents a historical version of a secret
type SecretVersion struct {
	ID        int64
	SecretID  string
	Version   int64
	Value     []byte
	CreatedAt time.Time
}
