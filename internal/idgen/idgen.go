package idgen

import (
	"crypto/rand"
	"encoding/base32"
)

var lowerBase32 = base32.NewEncoding("23456789abcdefghijkmnpqrstuvwxyz").WithPadding(base32.NoPadding)

// SecretID generates a unique secret identifier
func SecretID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return "sec-" + lowerBase32.EncodeToString(bytes), nil
}
