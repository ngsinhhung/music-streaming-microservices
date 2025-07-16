-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: IsUserExists :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);

-- name: CreateUser :one
INSERT INTO users (avatar, email, name, password)
VALUES ($1, $2, $3, $4)
RETURNING *;