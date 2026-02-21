package commands

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/fuzr0dah/locker/internal/db/migrations"
	"github.com/fuzr0dah/locker/internal/facade"
	"github.com/fuzr0dah/locker/internal/log"
	"github.com/fuzr0dah/locker/internal/server"
	"github.com/fuzr0dah/locker/internal/service"
	"github.com/fuzr0dah/locker/internal/storage"
	"github.com/fuzr0dah/locker/internal/storage/sqlite"

	"github.com/spf13/cobra"
)

var serverDescription = `A longer description that spans multiple lines and provides
comprehensive information about what your application does, how to use it,
and any important details users should know.`

func (f *CommandsFactory) NewServerCommand() *cobra.Command {
	var devMode bool
	cmd := &cobra.Command{
		Use:          "server",
		Short:        "A brief description of your application",
		Long:         serverDescription,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if devMode {
				f.print("server is running in dev mode")
			} else {
				return fmt.Errorf("production mode not yet implemented, use --dev flag")
			}

			auditFile, err := os.OpenFile(".build/audit.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
			if err != nil {
				return fmt.Errorf("open audit log: %w", err)
			}
			defer auditFile.Close()

			auditLogger := log.NewAuditLogger(auditFile)
			serverLogger := log.NewServerLogger(f.stdout)

			db, err := sqlite.OpenDB("")
			if err != nil {
				return fmt.Errorf("open db: %w", err)
			}

			if err := migrations.RunMigrations(db); err != nil {
				db.Close()
				return fmt.Errorf("run migrations: %w", err)
			}

			reader := sqlite.NewSecretReader(db)
			uowFactory := func() storage.UnitOfWork {
				return sqlite.NewUnitOfWork(db)
			}
			svc := service.NewSecretsService(reader, uowFactory, serverLogger)

			facadeLogger := serverLogger.With("component", "facade")
			facade := facade.NewFacade(svc, facadeLogger, auditLogger)

			srv, err := server.NewServer(facade, serverLogger)
			if err != nil {
				return fmt.Errorf("init server: %w", err)
			}

			ctx, cancel := context.WithCancel(cmd.Context())
			defer cancel()

			errChan := make(chan error, 1)
			go func() {
				errChan <- srv.Start()
			}()

			select {
			case err := <-errChan:
				return err
			case <-ctx.Done():
				shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer shutdownCancel()

				err := srv.Shutdown(shutdownCtx)
				if err != nil {
					return err
				}

				if err := db.Close(); err != nil {
					return fmt.Errorf("db close: %w", err)
				}

				return nil
			}
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&devMode, "dev", "d", false, "Run in dev mode")

	return cmd
}
