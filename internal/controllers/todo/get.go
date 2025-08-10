package todo

import (
	"context"

	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/patohru/todo-api/internal/database"
	"github.com/patohru/todo-api/internal/server/middleware"
)

func (r *TodoRoutes) GetTodoHandler(c fuego.ContextNoBody) (database.GetTodoRow, error) {
	ctx := context.Background()

	todo_id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return database.GetTodoRow{}, fuego.BadRequestError{
			Title: "Required UUID v4",
		}
	}
	account_id := c.Value(middleware.AuthorizationTokenKey).(uuid.UUID)

	queries := database.New(r.db)
	todo, err := queries.GetTodo(ctx, database.GetTodoParams{
		AccountID: account_id,
		ID: todo_id,
	})
	if err != nil {
		return database.GetTodoRow{}, fuego.BadRequestError{
			Title: "Cannot find todo with given id",
		}
	}

	return todo, nil
}
