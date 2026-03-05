-- name: CreateInitialVersion :one
INSERT INTO secret_versions (
    secret_id, version, value
) VALUES (
    ?1, 1, ?2
)
RETURNING *;

-- name: CreateNextVersion :one
INSERT INTO secret_versions (
    secret_id, version, value
) VALUES (
    ?1, ?2, ?3
)
RETURNING *;

-- name: GetSecretVersion :one
SELECT * FROM secret_versions
WHERE secret_id = ?1 AND version = ?2 LIMIT 1;

-- name: GetLastVersionForSecretId :one
SELECT * FROM secret_versions 
WHERE secret_id = ?1
ORDER BY version DESC
LIMIT 1;
