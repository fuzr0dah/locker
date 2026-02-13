-- name: GetSecretById :one
SELECT * FROM secrets
WHERE id = ?1 LIMIT 1;

-- name: GetSecretByName :one
SELECT * FROM secrets
WHERE name = ?1 LIMIT 1;

-- name: ListSecrets :many
SELECT * FROM secrets
ORDER BY name;

-- name: CreateSecret :one
INSERT INTO secrets (
    name, value
) VALUES (
    ?1, ?2
)
RETURNING *;
