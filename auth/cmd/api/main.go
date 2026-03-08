package main

import (
	"auth/config"
	"auth/internal/infrastructure/adapters/logger"
	"auth/internal/infrastructure/persistance/postgres"
	"log/slog"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	var cfg config.Config
	var db *gorm.DB

	cfg = config.MustLoad()
	logger.Setup(cfg.Env)

	db = postgres.MustOpenConnection(cfg.AuthDBServerConfig)
	_ = db

	slog.Debug("debug messages are enabled")
	slog.Info("Starting server...", slog.String("env", cfg.Env))

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
