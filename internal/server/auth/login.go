package auth

import (
	"context"
	"net/http"

	"github.com/patohru/todo-api/internal/database"
	"github.com/patohru/todo-api/internal/server/middleware"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Description Payload for /auth/login: user's email and password.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *AuthRoutes) LoginHandler(c *gin.Context) {
	ctx := context.Background()

	var request LoginRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, middleware.ApiError{
			Message: "Invalid login data",
		})
		return
	}

	queries := database.New(r.db)
	account, err := queries.GetAccountByEmail(ctx, request.Email)
	if err != nil {
		c.Error(&middleware.ApiError{
			Inner: err,
			Code: http.StatusBadRequest,
			Message: "Wrong email or password",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(request.Password)); err != nil {
		c.Error(&middleware.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Wrong email or password",
		})
		return
	}

	tokenString, err := r.jwtService.NewToken(account.ID.String())
	if err != nil {
		c.Error(&middleware.ApiError{
			Inner: err,
			Code: http.StatusInternalServerError,
			Message: "Failed to generate token",
		})
		return
	}

	c.String(http.StatusOK, tokenString)
}
