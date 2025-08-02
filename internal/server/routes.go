package server

import (
	"net/http"

	"github.com/patohru/todo-api/internal/handlers"
	"github.com/patohru/todo-api/internal/middleware"
	"github.com/gin-gonic/gin"

	_ "github.com/patohru/todo-api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(middleware.Cors())

	r.GET("/", handlers.Ping)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
