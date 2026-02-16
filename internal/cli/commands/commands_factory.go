package commands

import (
	"fmt"
	"io"
	"log/slog"
)

type CommandsFactory struct {
	stdout       io.Writer
	stderr       io.Writer
	stdin        io.Reader
	cliLogger    *slog.Logger
	serverLogger *slog.Logger
}

func NewCommandsFactory(stdout, stderr io.Writer, stdin io.Reader, cliLogger *slog.Logger, serverLogger *slog.Logger) *CommandsFactory {
	return &CommandsFactory{stdout: stdout, stderr: stderr, stdin: stdin, cliLogger: cliLogger, serverLogger: serverLogger}
}

func (f *CommandsFactory) print(format string, args ...any) {
	f.cliLogger.Info(format, args...)
	fmt.Fprintf(f.stdout, format+"\n", args...)
}

func (f *CommandsFactory) error(format string, args ...any) {
	f.cliLogger.Error(format, args...)
	fmt.Fprintf(f.stderr, "Error: "+format+"\n", args...)
}
