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
INSERT INTO addresses
(id, user_id, address, number, zip_code, city, state)
VALUES
($1, $2, $3, $4, $5, $6, $7);

-- name: GetAddresses :many
SELECT * FROM addresses WHERE user_id = $1;

-- name: GetAddressById :one
SELECT * FROM addresses WHERE id = $1 AND user_id = $2;

-- name: UpdateAddress :exec
UPDATE addresses
SET address = $3, number = $4, zip_code = $5, city = $6, state = $7
WHERE id = $1 AND user_id = $2;

-- name: DeleteAddress :exec
DELETE FROM addresses WHERE id = $1 AND user_id = $2;

-- name: GetProductById :one
SELECT * FROM products WHERE id = $1;

-- name: GetProducts :many
SELECT * FROM products;

-- name: CreateOrder :exec
INSERT INTO orders
(id, user_id, address_id, total_value, product_id, quantity)
VALUES
($1, $2, $3, $4, $5, $6);

-- name: GetUsersByProduct :many
SELECT u.name, u.email, o.total_value, o.quantity, o.id as order_id, a.state, a.address, a.number, a.zip_code, a.city
FROM users u
JOIN orders o
ON o.product_id = $1
JOIN addresses a
ON a.id = o.address_id
WHERE u.id = o.user_id;

