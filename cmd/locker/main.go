package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/fuzr0dah/locker/internal/cli/commands"
	"github.com/fuzr0dah/locker/internal/log"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logFile, err := os.OpenFile(".build/audit.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		logFile = os.Stderr
	}
	defer logFile.Close()

	cliLogger := log.NewCliLogger(logFile)
	serverLogger := log.NewServerLogger(os.Stdout)
	factory := commands.NewCommandsFactory(os.Stdout, os.Stderr, os.Stdin, cliLogger, serverLogger)
	cmd := factory.NewRootCommand()
	cmd.SetContext(ctx)

	if err := cmd.Execute(); err != nil {
		cliLogger.Warn("command failed", slog.Any("error", err))
		os.Exit(1)
	}
}
