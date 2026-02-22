package secrets

import (
	"crypto/rand"
	"encoding/base32"
	"regexp"
)

const Prefix = "sec-"
const Length = 24

var (
	lowerBase32 = base32.NewEncoding("23456789abcdefghijkmnpqrstuvwxyz").WithPadding(base32.NoPadding)
	validID     = regexp.MustCompile(`^` + Prefix + `[a-z0-9]{26}$`)
)

func SecretID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return "sec-" + lowerBase32.EncodeToString(bytes), nil
}

func IsValid(id string) bool {
	return validID.MatchString(id)
}
