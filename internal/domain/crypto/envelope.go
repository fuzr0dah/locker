package crypto

type Envelope interface {
	Seal(data []byte) ([]byte, error)
	Open(data []byte) ([]byte, error)
}
