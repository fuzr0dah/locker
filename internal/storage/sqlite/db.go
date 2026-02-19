package sqlite

import (
	"database/sql"
	"fmt"
)

// OpenDB opens a new SQLite database connection
// For in-memory mode: "file::memory:?cache=shared&_fk=on"
// For file mode: "file:/path/to/db.sqlite?_fk=on"
func OpenDB(dsn string) (*sql.DB, error) {
	if dsn == "" {
		dsn = "file::memory:?cache=shared&_fk=on&_journal_mode=WAL&_busy_timeout=5000"
	}

	conn, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	conn.SetMaxOpenConns(1)
	conn.SetMaxIdleConns(1)

	return conn, nil
}
