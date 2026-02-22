package commands

import (
	"fmt"
	"io"
)

type CommandsFactory struct {
	stdout io.Writer
	stderr io.Writer
	stdin  io.Reader
}

func NewCommandsFactory(stdout, stderr io.Writer, stdin io.Reader) *CommandsFactory {
	return &CommandsFactory{stdout: stdout, stderr: stderr, stdin: stdin}
}

func (f *CommandsFactory) print(format string, args ...any) {
	fmt.Fprintf(f.stdout, format+"\n", args...)
}

func (f *CommandsFactory) error(format string, args ...any) {
	fmt.Fprintf(f.stderr, "Error: "+format+"\n", args...)
}
