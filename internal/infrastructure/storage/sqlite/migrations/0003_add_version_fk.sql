-- +goose Up
ALTER TABLE secrets ADD COLUMN version_id INTEGER 
  REFERENCES secret_versions(id);
