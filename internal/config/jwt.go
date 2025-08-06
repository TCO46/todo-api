package config

type JwtConfig struct {
	Secret		[]byte	`env:"JWT_SECRET" envDefault:"secret"`
	JWT_REFRESH []byte	`env:"JWT_REFRESH"`
	ExpiredIn	int		`env:"JWT_EXPIRED_IN" envDefault:"24"`
}
