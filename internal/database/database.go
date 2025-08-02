package database

import (
	"context"
	"log"

	"github.com/patohru/todo-api/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

func Init() *pgxpool.Pool {
	cfg, _ := env.ParseAs[config.DatabaseConfig]()

	db, err := pgxpool.New(context.Background(), cfg.DatabaseURL())
	if err != nil {
		log.Fatal(err)
	}

	return db
}
