package config

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	LOCAL = "local"
	PROD  = "prod"
	DEV   = "dev"
)

type Config struct {
	Env              string `env:"ENV" env-default:"local"`
	HTTPServerConfig `yaml:",inline"`
	ServicesConfig   `yaml:",inline"`
	JWTConfig        `yaml:",inline"`
}

type JWTConfig struct {
	PublicKeyPath string `env:"JWT_PUBLIC_KEY_PATH" env-required:"true"`
	Issuer        string `env:"JWT_ISSUER"`
}

type ServicesConfig struct {
	AuthServiceURL string `env:"AUTH_SERVICE_URL" env-required:"true"`
}

type HTTPServerConfig struct {
	Address string `env:"HTTP_HOST" env-default:"0.0.0.0"`
	Port    string `env:"HTTP_PORT" env-default:"8000"`
	Cors    gin.HandlerFunc
}

func Init() Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using environment variables: %s", err)
	}

	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Error while reading environment variables: %s", err)
	}

	cfg.HTTPServerConfig.Cors = cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://localhost:3000", "http://localhost:5173", "http://words.skosyrsky.com", "http://127.0.0.1"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})

	return cfg
}
