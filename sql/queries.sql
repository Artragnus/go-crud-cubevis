-- name: CreateUser :exec
INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4);

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: UpdateUser :exec
UPDATE users set name = $2, email = $3, password = $4 WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: CreateAddress :exec
INSERT INTO addresses (id, user_id, address, number, zip_code, city, state) VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: GetAddresses :many
SELECT * FROM addresses WHERE user_id = $1;

-- name: GetAddressById :one
SELECT * FROM addresses WHERE id = $1 AND user_id = $2;
