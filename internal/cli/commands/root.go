package commands

import (
	"github.com/fuzr0dah/locker/internal/version"

	"github.com/spf13/cobra"
)

var rootDescription = `A longer description that spans multiple lines and provides
comprehensive information about what your application does, how to use it,
and any important details users should know.`

func (f *CommandsFactory) NewRootCommand() *cobra.Command {
	var (
		showVersion bool
	)

	cmd := &cobra.Command{
		Use:          "locker",
		Short:        "A brief description of your application",
		Long:         rootDescription,
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			if showVersion {
				f.print(version.GetVersion())
				return
			}
			f.print(rootDescription)
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&showVersion, "version", "v", false, "Print version information and quit")

	cmd.AddCommand(f.NewServerCommand())

	return cmd
}
