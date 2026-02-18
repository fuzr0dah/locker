-- name: CreateSecret :one
INSERT INTO secrets (
    id, name, version_id
) VALUES (
    ?1, ?2, NULL
)
RETURNING *;

-- name: InsertVersionIntoSecret :one
UPDATE secrets SET
    version_id = ?2
WHERE id = ?1 AND version_id IS NULL
RETURNING *;

-- name: GetSecretById :one
SELECT secrets.*, secret_versions.value FROM secrets
JOIN secret_versions ON secrets.version_id = secret_versions.id
WHERE secrets.id = ?1 LIMIT 1;

-- name: UpdateSecret :one
UPDATE secrets SET
    name = @name,
    version_id = @version_id,
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id AND version_id = @old_version_id
RETURNING *;

-- name: DeleteSecret :exec
DELETE FROM secrets
WHERE id = ?1;

-- name: ListSecrets :many
SELECT * FROM secrets
ORDER BY name;
