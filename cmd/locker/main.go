package main

import (
	"log/slog"
	"os"

	"github.com/fuzr0dah/locker/internal/commands"
)

func main() {
	factory := commands.NewCommandsFactory(os.Stdout, os.Stderr, os.Stdin)
	cmd := factory.NewRootCommand()

	if err := cmd.Execute(); err != nil {
		slog.Warn("command failed", slog.Any("error", err))
		os.Exit(1)
	}
}
