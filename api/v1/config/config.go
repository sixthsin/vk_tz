package config

import (
	"os"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
	Rest RestConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

type RestConfig struct {
	Port string
}

func LoadConfig() *Config {
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("AUTH_SERVICE_DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("JWT_SECRET"),
		},
		Rest: RestConfig{
			Port: os.Getenv("AUTH_SERVICE_REST_PORT"),
		},
	}
}
