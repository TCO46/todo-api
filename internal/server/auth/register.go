package auth

import (
	"context"
	"net/http"

	"github.com/patohru/todo-api/internal/database"
	"github.com/patohru/todo-api/internal/server/middleware"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Description Payload for /auth/register: user's email and password.
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *AuthRoutes) RegisterHandler(c *gin.Context) {
	ctx := context.Background()

	var request database.CreateAccountParams
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, middleware.ApiError{
			Message: "Invalid register data",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Error(&middleware.ApiError{
			Inner: err,
			Code: http.StatusBadGateway,
			Message: "Invalid register data",
		})
		return
	}
	request.Password = string(hashedPassword)

	queries := database.New(r.db) 
	id, err := queries.CreateAccount(ctx, request)
	if err != nil {
		c.Error(&middleware.ApiError{
			Inner: err,
			Code: http.StatusBadRequest,
			Message: "Account with given email already existed",
		})
		return
	}

	c.String(http.StatusOK, id.String())
}
