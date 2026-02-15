package commands

import (
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
