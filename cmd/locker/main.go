package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/fuzr0dah/locker/internal/cli/commands"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	factory := commands.NewCommandsFactory(os.Stdout, os.Stderr, os.Stdin)
	cmd := factory.NewRootCommand()
	cmd.SetContext(ctx)

	if err := cmd.Execute(); err != nil {
		slog.Warn("command failed", slog.Any("error", err))
		os.Exit(1)
	}
}
