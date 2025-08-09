package server

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/go-fuego/fuego"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"

	"github.com/patohru/todo-api/internal/server/ping"
	"github.com/patohru/todo-api/internal/server/todo"
	"github.com/patohru/todo-api/internal/server/auth"

	"github.com/patohru/todo-api/internal/config"
)

type Server struct {
	db *pgxpool.Pool
}

func NewServer() *fuego.Server {
	cfg, _ := env.ParseAs[config.ServerConfig]()

	s := fuego.NewServer(
		fuego.WithAddr(fmt.Sprintf(":%d", cfg.Port)),
	)

	auth.RegisterRoutes(s)
	ping.RegisterRoutes(s)
	todo.RegisterRoutes(s)

	return s
}
