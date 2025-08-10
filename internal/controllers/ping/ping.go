package ping

import (
	"github.com/go-fuego/fuego"
)

type PingResponse struct {
	Message string `json:"message"`
}

// PingExample	godoc
// @Summary		Ping back a pong
// @Description check if server is running
// @Produce json
// @Router / [get]
func PingHandler(c fuego.ContextNoBody) (*PingResponse, error) {
	return &PingResponse{
		Message: "Pong",
	}, nil
}

func RegisterRoutes(f *fuego.Server) {
	fuego.Get(f,  "/", PingHandler)
}
