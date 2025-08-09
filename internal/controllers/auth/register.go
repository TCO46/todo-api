package auth

import (
	"context"

	"github.com/go-fuego/fuego"
	"github.com/patohru/todo-api/internal/database"
	"golang.org/x/crypto/bcrypt"
)

// @Description Payload for /auth/register: user's email and password.
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *AuthRoutes) RegisterHandler(c fuego.ContextWithBody[database.CreateAccountParams]) (string, error) {

	var request database.CreateAccountParams
	request, err := c.Body()
	if err != nil {
		return "", fuego.BadRequestError{
			Title: "Invalid register data",
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fuego.BadRequestError{
			Title: "Invalid register data",
		}
	}
	request.Password = string(hashedPassword)

	queries := database.New(r.db) 
	ctx := context.Background()
	id, err := queries.CreateAccount(ctx, request)
	if err != nil {
		return "", fuego.BadRequestError{
			Title: "Account with given email already existed",
		}
	}

	return id.String(), nil
}
