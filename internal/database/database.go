package database

import (
	"fmt"
	"context"
	"log"
	"os"

	"github.com/patohru/todo-api/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
	"github.com/jackc/pgx/v5"
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
	pgxConfig, _ := pgxpool.ParseConfig(cfg.DatabaseURL())
	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
	    pgxUUID.Register(conn.TypeMap())
            return nil
	}

	instance, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
	    log.Fatal(err)
	    fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	if err = instance.Ping(ctx); err != nil {
	    fmt.Printf("Unable to ping database\n")
	}

	return instance 
}
