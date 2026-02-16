package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/fuzr0dah/locker/internal/server"

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

			srv, err := server.NewServer(f.serverLogger)
			if err != nil {
				f.cliLogger.Error("server initialization failed", "error", err)
				return fmt.Errorf("init server: %w", err)
			}

			f.cliLogger.Info("server started", "addr", ":8080")

			ctx, cancel := context.WithCancel(cmd.Context())
			defer cancel()

			errChan := make(chan error, 1)
			go func() {
				errChan <- srv.Start()
			}()

			select {
			case err := <-errChan:
				f.cliLogger.Error("server runtime error", "error", err)
				return err
			case <-ctx.Done():
				f.cliLogger.Info("server shutdown initiated", "reason", "signal received")

				shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer shutdownCancel()

				err := srv.Shutdown(shutdownCtx)
				if err != nil {
					f.cliLogger.Error("server shutdown failed", "error", err, "http_error", ctx.Err())
					return err
				}

				f.cliLogger.Info("server stopped gracefully", "shutdown_duration", "5s")
				return nil
			}
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&devMode, "dev", "d", false, "Run in dev mode")

	return cmd
}
