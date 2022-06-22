-- name: CreateAccount :one
INSERT INTO accounts (
  owner,
  balance,
  currency
) VALUES (
  $1, $2, $3
)RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 limit 1;

-- name: ListOfAccounts :many
SELECT * FROM accounts
ORDER BY id
limit $1 OFFSET $2;

-- name: UpdateAccount :one
update accounts 
set balance = $1
WHERE id=$2
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;
