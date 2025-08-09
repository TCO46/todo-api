package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/patohru/todo-api/internal/services/jwt"
)

const (
	authrizaion				string = "Authorization"
	bearer					string = "Bearer "
	AuthorizationTokenKey	string = "token"
)

func RequireAuthentication() gin.HandlerFunc {
	jwtService := jwt.New()

	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get(authrizaion)
		if authHeader == "" {
			c.Error(&ApiError{
				Code: http.StatusUnauthorized,
				Message: "Missing authorization header",
			})
			return
		}

		if !strings.HasPrefix(authHeader, bearer) {
		    c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be 'Bearer <token>'"})
		    return
		}
		tokenString := strings.TrimPrefix(authHeader, bearer)

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
