-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, api_key) 
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;


-- use "docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate" instead of sqlc generate you might need to run "docker pull kjconroy/sqlc"


-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key = $1;


-- name: GetUserByName :one
SELECT * FROM users where name = $1;