-- name: CreateUser :one
INSERT INTO users (id,name,username,email,password,mobile,role,is_active,created_at,updated_at,deleted_at)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) 
RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users;