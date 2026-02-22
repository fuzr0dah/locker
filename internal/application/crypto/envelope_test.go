package crypto

import (
	"encoding/binary"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/fuzr0dah/locker/internal/domain/crypto"
	"github.com/fuzr0dah/locker/internal/domain/crypto/mocks"
)

func TestEnvelopeService_Seal(t *testing.T) {
	t.Run("successfully creates envelope", func(t *testing.T) {
		// Arrange
		plaintext := []byte("sensitive data")
		encrypted := []byte("encrypted-by-cipher")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cipher := mocks.NewMockCipher(ctrl)
		svc := NewEnvelopeService(cipher)

		cipher.EXPECT().
			Encrypt(plaintext).
			Return(encrypted, nil).
			Times(1)

		// Act
		blob, err := svc.Seal(plaintext)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, blob)

		assert.Len(t, blob, 1+4+len(encrypted))
		assert.Equal(t, byte(0x01), blob[0], "version mismatch")
		assert.Equal(t, uint32(0), binary.BigEndian.Uint32(blob[1:5]), "keyID mismatch")
		assert.Equal(t, encrypted, blob[5:], "ciphertext mismatch")
	})

	t.Run("propagates encryption error", func(t *testing.T) {
		// Arrange
		plaintext := []byte("sensitive data")
		encryptErr := errors.New("cipher failed")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cipher := mocks.NewMockCipher(ctrl)
		svc := NewEnvelopeService(cipher)

		cipher.EXPECT().
			Encrypt(plaintext).
			Return(nil, encryptErr).
			Times(1)

		// Act
		blob, err := svc.Seal(plaintext)

		// Assert
		assert.ErrorIs(t, err, encryptErr)
		assert.Nil(t, blob)
	})
}

func TestEnvelopeService_Open(t *testing.T) {
	t.Run("validates envelope format", func(t *testing.T) {
		tests := []struct {
			name    string
			blob    []byte
			wantErr error
			errMsg  string
		}{
			{
				name:    "too short (< 5 bytes)",
				blob:    []byte{0x01, 0x00, 0x00},
				wantErr: crypto.ErrInvalidCiphertext,
				errMsg:  "invalid length check",
			},
			{
				name:    "wrong version",
				blob:    makeBlob(0x02, 0, []byte("data")),
				wantErr: crypto.ErrInvalidCiphertext,
				errMsg:  "version mismatch (expected 0x01, got 0x02)",
			},
			{
				name:    "wrong keyID",
				blob:    makeBlob(0x01, 999, []byte("data")),
				wantErr: crypto.ErrInvalidCiphertext,
				errMsg:  "keyID mismatch",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				svc := NewEnvelopeService(nil)
				got, err := svc.Open(tt.blob)
				assert.ErrorIs(t, err, tt.wantErr, tt.errMsg)
				assert.Nil(t, got)
			})
		}
	})

	t.Run("successfully opens envelope", func(t *testing.T) {
		// Arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cipher := mocks.NewMockCipher(ctrl)
		svc := NewEnvelopeService(cipher)

		ciphertext := []byte("encrypted-content")
		expectedPlain := []byte("hello world")
		blob := makeBlob(0x01, 0, ciphertext)

		cipher.EXPECT().
			Decrypt(ciphertext).
			Return(expectedPlain, nil).
			Times(1)

		// Act
		got, err := svc.Open(blob)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedPlain, got)
	})

	t.Run("propagates decryption error", func(t *testing.T) {
		// Arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cipher := mocks.NewMockCipher(ctrl)
		svc := NewEnvelopeService(cipher)

		ciphertext := []byte("corrupted")
		decryptErr := errors.New("integrity check failed")
		blob := makeBlob(0x01, 0, ciphertext)

		cipher.EXPECT().
			Decrypt(ciphertext).
			Return(nil, decryptErr).
			Times(1)

		// Act
		got, err := svc.Open(blob)

		// Assert
		assert.ErrorIs(t, err, decryptErr)
		assert.Nil(t, got)
	})
}

func TestEnvelopeService_RoundTrip(t *testing.T) {
	t.Run("seal then open returns original", func(t *testing.T) {
		// Arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cipher := mocks.NewMockCipher(ctrl)
		svc := NewEnvelopeService(cipher)

		original := []byte("top secret message")
		ciphertext := []byte("encrypted-bytes")

		cipher.EXPECT().
			Encrypt(original).
			Return(ciphertext, nil).
			Times(1)

		cipher.EXPECT().
			Decrypt(ciphertext).
			Return(original, nil).
			Times(1)

		// Act
		blob, err := svc.Seal(original)
		require.NoError(t, err)

		decrypted, err := svc.Open(blob)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, original, decrypted, "roundtrip failed")
	})
}

func makeBlob(version byte, keyID uint32, ciphertext []byte) []byte {
	blob := make([]byte, 1+4+len(ciphertext))
	blob[0] = version
	binary.BigEndian.PutUint32(blob[1:5], keyID)
	copy(blob[5:], ciphertext)
	return blob
}
