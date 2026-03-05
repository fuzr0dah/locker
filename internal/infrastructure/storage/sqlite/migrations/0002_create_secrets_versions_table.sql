-- +goose Up
CREATE TABLE secret_versions (
  id INTEGER PRIMARY KEY,
  secret_id TEXT NOT NULL,
  version INTEGER NOT NULL,
  value BLOB NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (secret_id) REFERENCES secrets(id) ON DELETE CASCADE,
  UNIQUE(secret_id, version)
);

CREATE INDEX idx_secret_versions_secret_id ON secret_versions(secret_id);
CREATE INDEX idx_secret_versions_created_at ON secret_versions(created_at);
