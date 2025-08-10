package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-fuego/fuego"
	"github.com/patohru/todo-api/internal/services/jwt"
)

const (
	authrizaion				string = "Authorization"
	bearer					string = "Bearer "
	AuthorizationTokenKey	string = "token"
)

func RequireAuthentication(next http.Handler) http.Handler {
	jwtService := jwt.New()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(authrizaion)
		if authHeader == "" {
			fuego.SendJSONError(w, nil, fuego.BadRequestError{
				Title: "Missing authorization header",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, bearer)

		token, err := jwtService.VerifyToken(tokenString)
		if err != nil {
			fuego.SendJSONError(w, nil, fuego.BadRequestError{
				Title: "Invalid authorization token",
			})
			return
		}

		ctx := context.WithValue(r.Context(), AuthorizationTokenKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
