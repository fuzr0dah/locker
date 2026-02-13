-- +goose Up
CREATE TABLE secrets (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  value BLOB NOT NULL,
  current_version INTEGER NOT NULL DEFAULT 1,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_secrets_name ON secrets(name);

-- +goose Down
DROP TABLE IF EXISTS secrets;
