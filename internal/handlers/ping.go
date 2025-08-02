package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingResponse struct {
	Message string `json:"message"`
}
// PingExample	godoc
// @Summary		Ping back a pong
// @Description check if server is running
// @Produce json
// @Router / [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, PingResponse{
		Message: "pong",
	})
}
