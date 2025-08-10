package todo

import (
	"context"

	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/moznion/go-optional"
	"github.com/patohru/todo-api/internal/database"
	"github.com/patohru/todo-api/internal/server/middleware"
)

// @Description Payload for /todo/create:
type UpdateRequest struct {
	Title		optional.Option[string]				`json:"title,omitempty"`
	Description	optional.Option[string]				`json:"description,omitempty"`
	Priority	optional.Option[database.Priority]	`json:"priority,omitempty"`
}

func (r *TodoRoutes) UpdateTodoHandler(c fuego.ContextWithBody[UpdateRequest]) (any, error) {
	ctx := context.Background()

	todo_id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Title: "Require UUID v4",
		}
	}

	request, err := c.Body()
	if err != nil {
		return nil, fuego.BadRequestError{
			Title: "Invalid update todo data",
		}
	}

	account_id := c.Value(middleware.AuthorizationTokenKey).(uuid.UUID)

	queries := database.New(r.db)
	if err := queries.UpdateTodo(ctx, database.UpdateTodoParams{
		AccountID: account_id,
		ID: todo_id,
		Title: request.Title,
	}); err != nil {
		return nil, fuego.BadRequestError{
			Title: "Cannot update with given id",
		}
	}

	return nil, nil
}
