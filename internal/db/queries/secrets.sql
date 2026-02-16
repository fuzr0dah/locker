-- name: CreateSecret :one
INSERT INTO secrets (
    id, name, value
) VALUES (
    ?1, ?2, ?3
)
RETURNING *;

-- name: GetSecretById :one
SELECT * FROM secrets
WHERE id = ?1 LIMIT 1;

-- name: UpdateSecret :one
UPDATE secrets SET
    name = ?2,
    value = ?3,
    current_version = current_version + 1,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?1 AND current_version = ?4
RETURNING *;

-- name: DeleteSecret :exec
DELETE FROM secrets
WHERE id = ?1;

-- name: ListSecrets :many
SELECT * FROM secrets
ORDER BY name;
