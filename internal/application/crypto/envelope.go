package crypto

import (
	"encoding/binary"

	"github.com/fuzr0dah/locker/internal/domain/crypto"
)

type envelopeService struct {
	cipher  crypto.Cipher
	version byte
	keyID   uint32
}

func NewEnvelopeService(cipher crypto.Cipher) *envelopeService {
	return &envelopeService{
		cipher:  cipher,
		version: 0x01,
		keyID:   0,
	}
}

func (e *envelopeService) Seal(data []byte) ([]byte, error) {
	ciphertext, err := e.cipher.Encrypt(data)
	if err != nil {
		return nil, err
	}
	blob := make([]byte, 1+4+len(ciphertext))
	blob[0] = e.version
	binary.BigEndian.PutUint32(blob[1:5], e.keyID)
	copy(blob[5:], ciphertext)
	return blob, nil
}

func (e *envelopeService) Open(data []byte) ([]byte, error) {
	if len(data) < 5 {
		return nil, crypto.ErrInvalidCiphertext
	}
	if data[0] != e.version {
		return nil, crypto.ErrInvalidCiphertext
	}
	if binary.BigEndian.Uint32(data[1:5]) != e.keyID {
		return nil, crypto.ErrInvalidCiphertext
	}
	plaintext, err := e.cipher.Decrypt(data[5:])
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
