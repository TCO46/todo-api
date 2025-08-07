package database

import (
	"fmt"
	"context"
	"log"
	"os"

	"github.com/patohru/todo-api/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

var instance *pgxpool.Pool

func NewPool() *pgxpool.Pool {
	ctx := context.Background()
	if instance != nil {
		return instance
	}

	cfg, _ := env.ParseAs[config.DatabaseConfig]()

	instance, err := pgxpool.New(ctx, cfg.DatabaseURL())
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	if err = instance.Ping(ctx); err != nil {
		fmt.Printf("Unable to ping database\n")
    }

	return instance 
}
