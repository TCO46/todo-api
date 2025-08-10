package todo

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/go-fuego/fuego"
	"github.com/patohru/todo-api/internal/database"
	"github.com/patohru/todo-api/internal/server/middleware"
)

type TodoRoutes struct {
	db				*pgxpool.Pool
}

func RegisterRoutes(f *fuego.Server) {
	r := TodoRoutes{
		db:				database.NewPool(),
	}

	todo := fuego.Group(f, "/todo")
	
	fuego.Use(todo, middleware.RequireAuthentication)

	fuego.Post(todo, "/create", r.CreateTodoHandler)
	fuego.Delete(todo, "/{id}", r.DeleteTodoHandler)
	fuego.Get(todo, "/{id}", r.GetTodoHandler)
	fuego.Patch(todo, "/{id}", r.UpdateTodoHandler)
}
