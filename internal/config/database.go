package config

import "fmt"

type DatabaseConfig struct {
	Database string `env:"DB_DATABASE"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Port     int    `env:"DB_PORT" envDefault:"5432"`
	Host     string `env:"DB_HOST" envDefault:"localhost"`
}

func (cfg *DatabaseConfig) DatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}
