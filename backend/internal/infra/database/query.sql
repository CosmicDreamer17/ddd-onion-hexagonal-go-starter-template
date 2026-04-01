-- name: CreateUser :exec
INSERT INTO users (id, email, password, created_at, updated_at)
VALUES (?, ?, ?, ?, ?);

-- name: GetUserByEmail :one
SELECT id, email, password, created_at, updated_at
FROM users
WHERE email = ?;

-- name: GetUserByID :one
SELECT id, email, password, created_at, updated_at
FROM users
WHERE id = ?;
