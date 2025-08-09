package todo

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (r *TodoRoutes) UpdateTodoHandler(c *gin.Context) {
	ctx := context.Background()

	todo_id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(&middleware.ApiError{
			Code: http.StatusBadRequest,
			Message: "Require UUID v4",
		})
		return
	}

	var request	UpdateRequest 
	if err := c.BindJSON(&request); err != nil {
		c.Error(&middleware.ApiError{
			Inner: err,
			Code: http.StatusBadRequest,
			Message: "Invalid update todo data",
		})
		return
	}

	account_id := c.MustGet(middleware.AuthorizationTokenKey).(uuid.UUID)

	queries := database.New(r.db)
	if err := queries.UpdateTodo(ctx, database.UpdateTodoParams{
		AccountID: account_id,
		ID: todo_id,
		Title: request.Title,
	}); err != nil {
		c.Error(&middleware.ApiError{
			Code: http.StatusBadRequest,
			Message: "Cannot update with given id",
		})
		return
	}

	c.Status(http.StatusOK)
}
