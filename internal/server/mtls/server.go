package mtls

import (
	"log/slog"

	"github.com/fuzr0dah/locker/internal/application/facade"
)

type server struct {
	logger  *slog.Logger
	facade  facade.SecretsFacade
	address string
}

func NewServer(f facade.SecretsFacade, logger *slog.Logger) (*server, error) {
	logger.Info("initializing mTLS server", "addr", "localhost:50051")

	return &server{
		facade:  f,
		logger:  logger,
		address: "localhost:8443",
	}, nil
}

func (s *server) Start() error {
	s.logger.Info("starting mTLS server", "address", s.address)
	return nil
}
