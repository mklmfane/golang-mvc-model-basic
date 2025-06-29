-- name: CreateAccount :one
INSERT INTO accounts_bank (full_name)
VALUES ($1)
RETURNING id, full_name, created_at;

-- name: GetAccount :one
SELECT id, full_name, created_at
FROM accounts_bank
WHERE id = $1;

-- name: UpdateAccount :one
UPDATE accounts_bank
SET full_name = $2
WHERE id = $1
RETURNING id, full_name, created_at;

-- name: DeleteAccount :exec
DELETE FROM accounts_bank
WHERE id = $1;
