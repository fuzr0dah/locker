package server

import (
	"fmt"
	"io"
	"os"

	appCrypto "github.com/fuzr0dah/locker/internal/application/crypto"
	"github.com/fuzr0dah/locker/internal/application/facade"
	"github.com/fuzr0dah/locker/internal/application/secrets"
	"github.com/fuzr0dah/locker/internal/domain/repository"
	infrCrypto "github.com/fuzr0dah/locker/internal/infrastructure/crypto"
	"github.com/fuzr0dah/locker/internal/infrastructure/log"
	"github.com/fuzr0dah/locker/internal/infrastructure/storage/sqlite"
	"github.com/fuzr0dah/locker/internal/infrastructure/storage/sqlite/db/migrations"
	httpsrv "github.com/fuzr0dah/locker/internal/server/http"
)

func NewRunnerWithDeps(stdout io.Writer, devMode bool) (*Runner, *Dependencies, error) {
	fmt.Fprintf(stdout, "starting server in dev mode: %v\n", devMode)

	auditFile, err := os.OpenFile(".build/audit.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return nil, nil, fmt.Errorf("open audit log: %w", err)
	}

	auditLogger := log.NewAuditLogger(auditFile)
	serverLogger := log.NewServerLogger(stdout)

	db, err := sqlite.OpenDB("")
	if err != nil {
		auditFile.Close()
		return nil, nil, fmt.Errorf("open db: %w", err)
	}

	if err := migrations.RunMigrations(db); err != nil {
		db.Close()
		auditFile.Close()
		return nil, nil, fmt.Errorf("run migrations: %w", err)
	}

	cipher := infrCrypto.NewAES()
	envelope := appCrypto.NewEnvelopeService(cipher)
	reader := sqlite.NewSecretReader(db)
	uowFactory := func() repository.UnitOfWork {
		return sqlite.NewUnitOfWork(db)
	}
	svc := secrets.NewSecretsService(envelope, reader, uowFactory, serverLogger)

	facadeLogger := serverLogger.With("component", "facade")
	f := facade.NewFacade(svc, facadeLogger, auditLogger)

	httpLogger := serverLogger.With("server", "http")
	httpSrv, err := httpsrv.NewServer(f, httpLogger)
	if err != nil {
		db.Close()
		auditFile.Close()
		return nil, nil, fmt.Errorf("init http server: %w", err)
	}

	// TODO: add mtls server

	runner := NewRunner(httpSrv)
	return runner, &Dependencies{DB: db, AuditFile: auditFile}, nil
}
