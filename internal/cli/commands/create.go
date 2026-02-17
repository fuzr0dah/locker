package commands

import (
	"fmt"

	"github.com/fuzr0dah/locker/internal/api"
	"github.com/fuzr0dah/locker/internal/cli/client"

	"github.com/spf13/cobra"
)

var createDescription = `A longer description that spans multiple lines and provides
comprehensive information about what your application does, how to use it,
and any important details users should know.`

func (f *CommandsFactory) NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "create",
		Short:        "A brief description of your application",
		Long:         createDescription,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := client.New("http://localhost:8080")
			secret, err := client.CreateSecret(cmd.Context(), &api.CreateSecretRequest{
				Name:  "my secret",
				Value: "{\"api_key\": \"my_secret_api_key\"}",
			})
			if err != nil {
				return fmt.Errorf("failed to create secret: %w", err)
			}
			f.print(secret.ID)
			return nil
		},
	}
	return cmd
}
