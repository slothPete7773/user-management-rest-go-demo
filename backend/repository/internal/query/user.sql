-- name: GetAccountInfoByEmail :one
SELECT * FROM account AS a
JOIN account_information AS ai
  ON a.id = ai.account_id
WHERE email = $1 LIMIT 1;


-- name: GetAccount :one
SELECT * FROM account AS a
JOIN account_information AS ai
  ON a.id = ai.account_id
WHERE username = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM account AS a
JOIN account_information AS ai
  ON a.id = ai.account_id
ORDER BY username;

-- name: InsertAccount :one
INSERT INTO account (
  id, username, email, created_at, updated_at
)
VALUES (
  $1, $2, $3, NOW(), NOW()
)
RETURNING *;

-- name: UpdateAccount :one
UPDATE account
  set
    username = $2,
    email = $3
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
UPDATE account
SET deleted_at = NOW()
WHERE id = $1;

-- name: InsertAccountInfo :one
INSERT INTO account_information (
  id
  , display_name
  , favorite_number
  , homeworld_realm
  , account_id
  , created_at
  , updated_at
) VALUES (
  $1, $2, $3, $4, $5, NOW(), NOW()
)
RETURNING *;

-- name: UpdateAccountInfo :one
UPDATE account_information
SET
  display_name = $1
  , favorite_number = $2
  , homeworld_realm = $3
  , updated_at = NOW()
WHERE account_id = $4
RETURNING *;

-- name: DeleteAccountInfo :exec
UPDATE account_information
SET deleted_at = NOW()
WHERE account_id = $1;