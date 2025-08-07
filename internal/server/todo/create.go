package todo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (r *TodoRoutes) CreateTodoHandler(c *gin.Context) {
	ctx := context.Background()

	var request CreateRequest
	if err := c.BindJSON(&request); err != nil {
		c.Error(&middleware.ApiError{
			Inner: err,
			Code: http.StatusBadRequest,
			Message: "Invalid create todo data",
		})
		return
	}

	account_id := c.MustGet(middleware.AuthorizationTokenKey).(uuid.UUID)
	fmt.Print(account_id)

	queries := database.New(r.db)
	id, err := queries.CreateTodo(ctx, database.CreateTodoParams{
		AccountID: account_id,
		Title: request.Title,
		Description: request.Description,
		Priority: request.Priority,
	})
	if err != nil {
		c.Error(&middleware.ApiError{
			Inner: err,
			Code: http.StatusBadRequest,
			Message: "Title already exist",
		})
		return
	}

	c.String(http.StatusOK, id.String())
}
