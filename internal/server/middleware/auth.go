package middleware

import (
	"net/http"
	"strings"

	"github.com/patohru/todo-api/internal/services/jwt"
	"github.com/gin-gonic/gin"
)

const (
	authrizaion				string = "Authorization"
	bearer					string = "Bearer: "
	AuthorizationTokenKey	string = "token"
)

func RequireAuthentication() gin.HandlerFunc {
	jwtService := jwt.New()

	return func(c *gin.Context) {
		authHeader := c.GetHeader(authrizaion)
		if authHeader == "" {
			c.Error(&ApiError{
				Code: http.StatusUnauthorized,
				Message: "Missing authorization header",
			})
			return
		}

		tokenString, isBearer := strings.CutPrefix(authHeader, bearer)
		if !isBearer {
			c.Error(&ApiError{
				Code: http.StatusUnauthorized,
				Message: "Missing authorization token",
			})
			return
		}

		token, err := jwtService.VerifyToken(tokenString)
		if err != nil {
			c.Error(&ApiError{
				Inner: err,
				Code: http.StatusForbidden,
				Message: "Invalid authorization token",
			})
			return	
		}

		c.Set(AuthorizationTokenKey, token)
	}
}
