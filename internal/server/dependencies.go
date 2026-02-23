package server

import (
	"database/sql"
	"fmt"
	"os"
)

type Dependencies struct {
	DB        *sql.DB
	AuditFile *os.File
}

func (d *Dependencies) Close() error {
	if d.DB != nil {
		if err := d.DB.Close(); err != nil {
			return fmt.Errorf("close db: %w", err)
		}
	}
	if d.AuditFile != nil {
		if err := d.AuditFile.Close(); err != nil {
			return fmt.Errorf("close audit file: %w", err)
		}
	}
	return nil
}
