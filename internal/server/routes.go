package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/patohru/todo-api/internal/server/ping"
	"github.com/patohru/todo-api/internal/server/auth"
	"github.com/patohru/todo-api/internal/server/middleware"

	_ "github.com/patohru/todo-api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) RegisterRoutes() http.Handler {
	g := gin.Default()

	g.Use(middleware.Cors())
	g.Use(middleware.ErrorHandler())

	auth.RegisterRoutes(g)
	ping.RegisterRoutes(g)

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return g 
}
