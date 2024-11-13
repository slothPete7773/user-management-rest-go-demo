-- name: GetAccount :one
SELECT * FROM account
WHERE username = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM account
ORDER BY username;

-- name: InsertAccount :one
INSERT INTO account (
  username, email
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateAccount :one
UPDATE account
  set email = $2
WHERE username = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE username = $1;

-- name: InsertAccountInfo :one
INSERT INTO account_information (
  id, display_name, favorite_number, homeworld_realm, account_id
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateAccountInfo :one

-- name: GetAccountInfoByUsername :one
