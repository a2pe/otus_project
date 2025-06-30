package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Host   string
	Port   int
	Secret string
}

type AuthConfig struct {
	AdminLogin    string
	AdminPassword string
	JWTSecret     string
}

var Cfg Config

func NewConfig() *Config {
	Cfg = Config{
		Host:   "0.0.0.0",
		Port:   8080,
		Secret: "secret",
	}
	return &Cfg
}

func LoadAuthConfig() AuthConfig {
	_ = godotenv.Load()

	login := os.Getenv("ADMIN_LOGIN")
	pass := os.Getenv("ADMIN_PASSWORD")
	secret := os.Getenv("JWT_SECRET")

	if login == "" || pass == "" || secret == "" {
		log.Fatal("Missing one or more required auth env variables")
	}

	return AuthConfig{
		AdminLogin:    login,
		AdminPassword: pass,
		JWTSecret:     secret,
	}
}
