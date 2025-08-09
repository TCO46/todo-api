package auth

import(
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/go-fuego/fuego"

	"github.com/patohru/todo-api/internal/database"
	"github.com/patohru/todo-api/internal/services/jwt"
)

type AuthRoutes struct {
	db				*pgxpool.Pool
	jwtService		*jwt.JwtService
}

func RegisterRoutes(f *fuego.Server) {
	r := AuthRoutes{
		db:				database.NewPool(),
		jwtService:		jwt.New(),
	}

	fuego.Post(f, "/auth/login", r.LoginHandler)
	fuego.Post(f, "/auth/register", r.RegisterHandler)
}
