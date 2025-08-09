package middleware

import (
	"github.com/patohru/todo-api/internal/config"

	"github.com/caarlos0/env/v11"
)

func Cors() gin.HandlerFunc {
	cfg, _ := env.ParseAs[config.CorsConfig]()

	return cors.New(cors.Config{
		AllowOrigins:     []string{cfg.AllowOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	})
}
