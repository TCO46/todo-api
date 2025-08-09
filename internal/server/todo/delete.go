package todo

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/patohru/todo-api/internal/database"
	"github.com/patohru/todo-api/internal/server/middleware"
)

func (r *TodoRoutes) DeleteTodoHandler(c *gin.Context) {
	ctx := context.Background()

	todo_id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(&middleware.ApiError{
			Code: http.StatusBadRequest,
			Message: "Require UUID v4",
		})
		return
	}
	account_id := c.MustGet(middleware.AuthorizationTokenKey).(uuid.UUID)
	queries := database.New(r.db)
	if err := queries.DeleteTodo(ctx, database.DeleteTodoParams{
		AccountID: account_id,
		ID: todo_id,
	}); err != nil {
		c.Error(&middleware.ApiError{
			Code: http.StatusBadRequest,
			Message: "Cannot file delete with given ID",
		})
	}

	c.Status(http.StatusOK)
}
