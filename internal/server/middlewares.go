package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func loggerMiddleware(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ctx := r.Context()

			logger.Info("started http request",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("user_agent", r.UserAgent()),
			)

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			next.ServeHTTP(ww, r.WithContext(ctx))

			logger.Info("completed http request",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("user_agent", r.UserAgent()),
				slog.Int("status", ww.Status()),
				slog.Int64("duration", time.Since(start).Milliseconds()),
			)
		})
	}
}
