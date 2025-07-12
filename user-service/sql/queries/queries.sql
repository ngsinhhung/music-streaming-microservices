-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: IsUserExists :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);