package config

import (
	"time"
)

type Config struct {
	Server   ServerConfig
	Auth     AuthConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type AuthConfig struct {
	JWTSecret     string
	TokenDuration time.Duration
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func LoadConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port: ":8080",
		},
		Auth: AuthConfig{
			JWTSecret:     "your-secret-key-change-in-production",
			TokenDuration: 24 * time.Hour,
		},
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "user",
			Password: "password",
			Database: "dbname",
		},
	}
}
