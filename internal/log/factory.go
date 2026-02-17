package log

import (
	"io"
	"log/slog"
	"time"

	"github.com/lmittmann/tint"
)

func NewCliLogger(stdout io.Writer) *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	return logger
}

// TODO сreate separate logger function for debug and production.
func NewServerLogger(stdout io.Writer) *slog.Logger {
	return slog.New(tint.NewHandler(stdout, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.TimeOnly,
		AddSource:  true,
	}))
}
