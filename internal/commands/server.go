package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/fuzr0dah/locker/internal/crypto"
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
				fmt.Fprintf(f.stdout, "Running in dev mode\n")
			} else {
				return fmt.Errorf("production mode not yet implemented, use --dev flag")
			}

			fmt.Fprintf(f.stdout, "TODO: Master key - %s\n", crypto.GenerateMasterKey())

			srv, err := server.NewServer(f.stdout)
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
				return srv.Shutdown(shutdownCtx)
			}
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&devMode, "dev", "d", false, "Run in dev mode")

	return cmd
}
