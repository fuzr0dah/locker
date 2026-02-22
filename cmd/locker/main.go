package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fuzr0dah/locker/internal/commands"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	factory := commands.NewCommandsFactory(os.Stdout, os.Stderr, os.Stdin)
	cmd := factory.NewRootCommand()
	cmd.SetContext(ctx)

	if err := cmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
