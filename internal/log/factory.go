package log

import (
	"io"
	"log/slog"
)

func NewCliLogger(stdout io.Writer) *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	return logger
}

func NewServerLogger(stdout io.Writer) *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	return logger
}
