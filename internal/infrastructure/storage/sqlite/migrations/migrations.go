package migrations

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var migrationsFS embed.FS

// RunMigrations executes all pending migrations on the given SQLite database
func RunMigrations(db *sql.DB) error {
	// Create goose provider with embedded filesystem
	provider, err := goose.NewProvider(
		goose.DialectSQLite3,
		db,
		migrationsFS,
	)
	if err != nil {
		return fmt.Errorf("create goose provider: %w", err)
	}

	// Run all pending migrations
	if _, err := provider.Up(context.Background()); err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}

	return nil
}
