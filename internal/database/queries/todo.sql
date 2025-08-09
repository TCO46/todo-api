-- name: CreateTodo :one
INSERT INTO todos(account_id, title, description, priority)
VALUES (@account_id, @title, @description, @priority)
RETURNING id;

-- name: GetTodo :one
SELECT title, description, priority, is_done, created_at
FROM todos
WHERE id = @id AND account_id = @account_id;

-- name: GetTodos :many
SELECT id, title, description, priority, is_done, created_at
FROM todos
WHERE account_id = @account_id AND
    (
        @query IS NULL OR
        title LIKE '%' || @query || '%' OR
        description LIKE '%' || @query || '%'
    ) AND
    (@priority IS NULL OR priority = @priority) AND
    (@is_done IS NULL OR is_done = @is_done);

-- name: UpdateTodo :exec
UPDATE todos
SET
    title = COALESCE(sqlc.narg('title'), title),
    description = COALESCE(sqlc.narg('description'), content),
    priority = COALESCE(sqlc.narg('priority'), priority)
WHERE id = @id AND account_id = @account_id AND is_done = false;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = @id AND account_id = @account_id;
