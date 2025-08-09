package auth

import (
	"context"

	"github.com/go-fuego/fuego"
	"github.com/patohru/todo-api/internal/database"
	"golang.org/x/crypto/bcrypt"
)

// @Description Payload for /auth/login: user's email and password.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *AuthRoutes) LoginHandler(c fuego.ContextWithBody[LoginRequest]) (string, error) {
	request, err := c.Body()
	if err != nil {
		return "", fuego.BadRequestError{
			Title: "Invalid login data",
		}
	}

	queries := database.New(r.db)
	ctx := context.Background()

	account, err := queries.GetAccountByEmail(ctx, request.Email)
	if err != nil {
		return "", fuego.BadRequestError{
			Title: "Wrong email or password",
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(request.Password)); err != nil {
		return "", fuego.BadRequestError{
			Title: "Wrong email or password",
		}
	}

	tokenString, err := r.jwtService.NewToken(account.ID.String())
	if err != nil {
		return "", fuego.BadRequestError{
			Title: "Failed to generate token",
		}
	}

	return tokenString, nil
}
