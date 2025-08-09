package todo

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/patohru/todo-api/internal/database"
	"github.com/patohru/todo-api/internal/server/middleware"
)

type TodoRoutes struct {
	db				*pgxpool.Pool
}

func RegisterRoutes(g *gin.Engine) {
	r := TodoRoutes{
		db:				database.NewPool(),
	}

	todo := g.Group("/").Use(middleware.RequireAuthentication())
	todo.POST("/todo/create", r.CreateTodoHandler)
	todo.DELETE("/todo/:id", r.DeleteTodoHandler)
	todo.GET("/todo/:id", r.GetTodoHandler)
	todo.PATCH("/todo/:id", r.UpdateTodoHandler)
}
