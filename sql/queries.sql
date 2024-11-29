-- name: CreateUser :exec
INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4);

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users set name = $2, email = $3, password = $4 WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
