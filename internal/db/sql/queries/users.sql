
-- User---------------------------------------------------------------

-- name: CreateUser :one
INSERT INTO users
    (id,name,username,email,password,mobile,role,is_active,max_sessions,created_at,updated_at,deleted_at)
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;

-- name: GetAllUsers :many
SELECT *
FROM users;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id=$1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email=$1;




-- OTP----------------------------------------------------------------

-- name: CreateOTP :one 
INSERT INTO user_otps
    (user_id,otp_code,expires_at,is_used,created_at)
-- id is created internally
VALUES
    ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetLatestOTPFromUser :one
SELECT *
FROM user_otps
WHERE user_id=$1
ORDER BY created_at DESC LIMIT 1;

-- name: DeleteUserOTPByUserId :exec
DELETE FROM user_otps
WHERE user_id=$1;

-- name: MarkOTPAsUsed :exec
UPDATE user_otps SET is_used = TRUE WHERE user_id=$1;



-- Session----------------------------------------------------------

-- name: CreateUserSession :one
INSERT INTO user_sessions
    (session_id,user_id,created_at,expires_at,ip_address)
VALUES
    ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserSessionByUserId :one
SELECT *
FROM user_sessions
WHERE session_id=$1
ORDER BY created_at DESC LIMIT 1;

-- name: GetUserSessionById :one
SELECT *
FROM user_sessions
WHERE user_id=$1
ORDER BY created_at DESC LIMIT 1;

-- name: DeleteUserSession :exec 
DELETE FROM user_sessions
WHERE session_id=$1;




--Refresh Token --------------------------------------------------------


-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens
    (id, user_id,session_id, token_hash, created_at, expires_at, revoked_at)
VALUES
    ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetRefreshToken :one
SELECT *
FROM refresh_tokens
WHERE session_id = $1
ORDER BY created_at DESC
LIMIT 1;


-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens
SET revoked_at = NOW() 
WHERE token_hash = $1;


-- name: RotateRefreshToken :exec
UPDATE refresh_tokens
SET revoked_at = NOW(), replaced_by = $2
WHERE id = $1;
