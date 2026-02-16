-- name: InsertSecretVersion :one
INSERT INTO secret_versions (
    secret_id, version, value
) VALUES (
    ?1, ?2, ?3
)
RETURNING *;
