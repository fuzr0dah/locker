package crypto

import "errors"

var (
	ErrDecryptionFailed  = errors.New("decryption failed: data corrupted or invalid key")
	ErrInvalidCiphertext = errors.New("ciphertext format invalid")
)
