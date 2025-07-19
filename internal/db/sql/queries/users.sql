-- name: CreateUser :one
INSERT INTO users (id,name,username,email,password,mobile,role,is_active,created_at,updated_at,deleted_at)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) 
RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: CreateOTP :one 
INSERT INTO user_otps (user_id,otp_code,expires_at,is_used,created_at)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: GetLatestOTPFromUser :one
SELECT * FROM user_otps WHERE user_id=$1 ORDER BY created_at DESC LIMIT 1;

-- name: DeleteUserOTPByUserId :exec
DELETE FROM user_otps WHERE user_id=$1;

-- name: CreateSession :one
INSERT INTO user_sessions (session_id,user_id,created_at,expires_at,ip_address)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;