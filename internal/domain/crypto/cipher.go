package crypto

//go:generate mockgen -source=cipher.go -destination=mocks/mock_cipher.go -package=mocks
type Cipher interface {
	Encrypt(data []byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
}
