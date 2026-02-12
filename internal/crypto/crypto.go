package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateMasterKey() string {
	hash := sha256.Sum256([]byte("random string"))
	return hex.EncodeToString(hash[:])
}
