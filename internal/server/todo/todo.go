package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/patohru/todo-api/internal/database"
)

type TodoRoutes struct {
	db				*pgxpool.Pool
}

func RegisterRoutes(g *gin.Engine) {
	r := TodoRoutes{
		db:				database.NewPool(),
	}

	g.POST("/todo/create", middleware.RequireAuthentication(), r.CreateTodoHandler)
}
