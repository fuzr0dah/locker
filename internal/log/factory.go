package log

import (
	"io"
	"log/slog"
	"time"

	"github.com/lmittmann/tint"
)

func NewAuditLogger(w io.Writer) *slog.Logger {
	return slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "value" || a.Key == "password" || a.Key == "master_key" {
				return slog.Attr{}
			}
			return a
		},
	}))
}

// TODO сreate separate logger function for debug and production.
func NewServerLogger(stdout io.Writer) *slog.Logger {
	return slog.New(tint.NewHandler(stdout, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.TimeOnly,
		AddSource:  true,
	}))
}
