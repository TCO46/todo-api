package todo

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/patohru/todo-api/internal/database"
	"github.com/patohru/todo-api/internal/server/middleware"
)

func (r *TodoRoutes) GetTodoHandler(c *gin.Context) {
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
	todo, err := queries.GetTodo(ctx, database.GetTodoParams{
		AccountID: account_id,
		ID: todo_id,
	})
	if err != nil {
		c.Error(&middleware.ApiError{
			Code: http.StatusBadRequest,
			Message: "Cannot find todo with given id",
		})
		return
	}

	c.JSON(http.StatusOK, todo)

}
