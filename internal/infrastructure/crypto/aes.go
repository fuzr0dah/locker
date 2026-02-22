package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/fuzr0dah/locker/internal/domain/crypto"
)

const password = "random string"

func key() []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:]
}

type AES struct {
	key []byte
}

func NewAES() *AES {
	return &AES{
		key: key(),
	}
}

func (a *AES) Encrypt(data []byte) ([]byte, error) {
	nonce := make([]byte, 12)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, data, nil)

	result := make([]byte, 12+len(ciphertext))
	copy(result[0:12], nonce)
	copy(result[12:], ciphertext)

	return result, nil
}

func (a *AES) Decrypt(data []byte) ([]byte, error) {
	if len(data) < 12+16 { // 12 + 16 = nonce + tag
		return nil, crypto.ErrInvalidCiphertext
	}

	nonce := data[0:12]
	ciphertext := data[12:]

	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", crypto.ErrDecryptionFailed, err)
	}
	return plaintext, nil
}

func GenerateMasterKey() string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
