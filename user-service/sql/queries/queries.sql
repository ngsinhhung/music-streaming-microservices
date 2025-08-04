-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: IsUserExists :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);

-- name: CreateUser :one
INSERT INTO users (avatar, email, name, password)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetLoginSessionByUserId :one
SELECT * FROM user_login_session WHERE user_id = $1 LIMIT 1;

-- name: CreateUserLoginSession :one
INSERT INTO user_login_session (uuid, user_id, public_key, rf_token, rf_token_used)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateLoginSession :one
