package todo

import (
	"context"

	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/patohru/todo-api/internal/database"
	"github.com/patohru/todo-api/internal/server/middleware"
)

func (r *TodoRoutes) DeleteTodoHandler(c fuego.ContextNoBody) (any, error) {

	todo_id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Title: "Require UUID v4",
		}
	}
	account_id := c.Value(middleware.AuthorizationTokenKey).(uuid.UUID)
	ctx := context.Background()
	queries := database.New(r.db)
	if err := queries.DeleteTodo(ctx, database.DeleteTodoParams{
		AccountID: account_id,
		ID: todo_id,
	}); err != nil {
		return nil, fuego.BadRequestError{
			Title: "Cannot find delete with given ID",
		}
	}

	return nil, nil
}
