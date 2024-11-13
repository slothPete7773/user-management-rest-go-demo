-- name: CreateProject :one
INSERT INTO project (id, title, sector, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING *;

-- name: UpdateProject :one
UPDATE project
SET title = $1, sector = $2, updated_at = NOW()
WHERE id = $3
RETURNING *;

-- name: DeleteProject :exec
UPDATE project
SET deleted_at = NOW()
WHERE id = $1;

-- name: CreateProjectAssignment :one
INSERT INTO project_account_assignment (id, project_id, account_id, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING *;

-- name: DeleteProjectAssignment :exec
UPDATE project_account_assignment
SET deleted_at = NOW()
WHERE id = $1;

-- name: GetProjects :many
SELECT *
FROM project
WHERE deleted_at IS NULL;

-- name: GetAssignedProjectsByUserEmail :many
SELECT
  p.*
FROM project p
JOIN project_account_assignment paa
    ON p.id = paa.project_id
JOIN account a
    ON paa.account_id = a.id
WHERE a.email = $1 AND p.deleted_at IS NULL;

-- name: GetProjectMembersFromProjectId :many
SELECT
  a.email
  , a.username
FROM project_account_assignment paa
JOIN account a
    ON paa.account_id = a.id
WHERE paa.project_id = $1;