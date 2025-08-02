package server

import (
	"net/http"

	"github.com/patohru/todo-api/internal/handlers"
	"github.com/patohru/todo-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(middleware.Cors())

	r.GET("/", handlers.Ping)

	return r
}
