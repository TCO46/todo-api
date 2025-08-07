package auth

import(
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/patohru/todo-api/internal/database"
	"github.com/patohru/todo-api/internal/services/jwt"
)

type AuthRoutes struct {
	db				*pgxpool.Pool
	jwtService		*jwt.JwtService
}

func RegisterRoutes(g *gin.Engine) {
	r := AuthRoutes{
		db:				database.NewPool(),
		jwtService:		jwt.New(),
	}

	g.POST("/auth/login", r.LoginHandler)
	g.POST("/auth/register", r.RegisterHandler)
}
