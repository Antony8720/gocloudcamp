package database

import (
	"os"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Name     string
	Password string
}

// Configure database connection from env
func NewConfig() *Config {
	cfg := Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Name:     os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	return &cfg
}
