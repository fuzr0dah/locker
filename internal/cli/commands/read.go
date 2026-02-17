package commands

import (
	"encoding/json"
	"fmt"

	"github.com/fuzr0dah/locker/internal/cli/client"

	"github.com/spf13/cobra"
)

var readDescription = `A longer description that spans multiple lines and provides
comprehensive information about what your application does, how to use it,
and any important details users should know.`

func (f *CommandsFactory) NewReadCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "read",
		Short:        "A brief description of your application",
		Long:         readDescription,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := client.New("http://localhost:8080")
			if len(args) == 0 {
				return fmt.Errorf("secret id is required")
			}
			secret, err := client.GetSecret(cmd.Context(), args[0])
			if err != nil {
				return fmt.Errorf("failed to read secret: %w", err)
			}
			data, err := json.MarshalIndent(secret, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal secret: %w", err)
			}
			f.print("%s", string(data))
			return nil
		},
	}
	return cmd
}
