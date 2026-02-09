package commands

import (
	"fmt"

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
		Run: func(cmd *cobra.Command, args []string) {
			if devMode {
				fmt.Println("Hello, Dev Server!")
				return
			}
			fmt.Println("Hello, Server!")
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&devMode, "dev", "d", false, "Run in dev mode")

	return cmd
}
