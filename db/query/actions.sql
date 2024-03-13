-- name: Createusertable :one
INSERT INTO users (
  first_name,
  last_name,
  user_name,
  email,
  password
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: Getuser :one
SELECT * FROM users 
WHERE user_id = $1;

-- name: Listuser :many
SELECT * FROM users 
ORDER BY user_id
LIMIT $1
OFFSET $2;

-- name: Updateuser :one
UPDATE users 
SET first_name = $2, last_name = $3, user_name = $4
WHERE user_id = $1
RETURNING *;

-- name: Deleteuser :exec
DELETE FROM users 
WHERE user_id = $1;

-- name: Createtodos :one
INSERT INTO todos(
  title,
  time,
  date
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: Gettodos :one
SELECT * FROM todos
WHERE todo_id = $1
ORDER BY todo_id;

-- name: Listtodos :many
SELECT * FROM todos
WHERE todo_id = $1
ORDER BY todo_id;

-- name: Updatetodos :one
UPDATE todos
SET title = $2, time = $3, date = $4
WHERE todo_id = $1
RETURNING *;

-- name: Deletetodos :exec
DELETE FROM todos 
WHERE todo_id = $1;

