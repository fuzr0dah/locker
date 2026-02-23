package commands

import (
	"fmt"

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
			runner, deps, err := server.NewRunnerWithDeps(f.stdout, devMode)
			if err != nil {
				return err
			}
			defer deps.Close()

			runner.Start()
			if err := runner.Wait(); err != nil {
				return fmt.Errorf("server error: %w", err)
			}
			return nil
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&devMode, "dev", "d", false, "Run in dev mode")

	return cmd
}
