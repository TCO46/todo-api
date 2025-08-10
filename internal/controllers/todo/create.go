package todo

import (
	"context"

	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/patohru/todo-api/internal/database"
	"github.com/patohru/todo-api/internal/server/middleware"
)

// @Description Payload for /todo/create:
type CreateRequest struct {
	Title		string
	Description string
	Priority	database.Priority
}

func (r *TodoRoutes) CreateTodoHandler(c fuego.ContextWithBody[CreateRequest]) (string, error) {

	request, err := c.Body()
	if err != nil {
		return "", fuego.BadRequestError{
			Title: "Invalid create todo data",
		}
	}

	account_id := c.Value(middleware.AuthorizationTokenKey).(uuid.UUID)

	queries := database.New(r.db)
	ctx := context.Background()
	id, err := queries.CreateTodo(ctx, database.CreateTodoParams{
		AccountID: account_id,
		Title: request.Title,
		Description: request.Description,
		Priority: request.Priority,
	})
	if err != nil {
		return "", fuego.BadRequestError{
			Title: "Title already exist",
		}
	}

	return id.String(), nil
}
