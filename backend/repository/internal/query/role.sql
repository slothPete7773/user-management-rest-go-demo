-- name: CreateRole :one
INSERT INTO user_role (id, role_title, created_at, updated_at)
VALUES ($1, $2, NOW(), NOW())
RETURNING *;

-- name: UpdateRole :one
UPDATE user_role
SET
    role_title = $1,
    updated_at = NOW()
WHERE id = $2
RETURNING *;

-- name: DeleteRole :exec
UPDATE user_role
SET
    deleted_at = NOW()
WHERE id = $1;

-- name: AssignRole :one
INSERT INTO account_role (id, account_id, user_role_id, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING *;

-- name: GetRoleByEmail :one
SELECT
    a.id
    , a.email
    , ur.role_title
FROM user_role AS ur
JOIN account_role AS ar
    ON ur.id = ar.user_role_id
JOIN account AS a
    ON ar.account_id = a.id
WHERE a.email = $1
LIMIT 1;