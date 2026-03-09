package main

import (
	"auth/config"
	"auth/internal/application/usecase"
	"auth/internal/infrastructure/postgres/repositories"
	httpv1 "auth/pkg/http"
	"auth/pkg/logger"
	"auth/pkg/postgres"
	"auth/pkg/router"
	"context"
	"log/slog"
	"os"
	"time"
)

func main() {

	cfg := config.Init()

	logger := logger.Init(cfg.Env)

	db := postgres.Init(cfg.PostgresConfig)

	logger.Debug("debug messages are enabled")
	logger.Info("Starting server...", slog.String("env", cfg.Env))

	// Initialize repositories
	userRepository := repositories.NewUserRepository(db)

	// Initialize usecases
	usecases := usecase.NewUserUseCases(userRepository, logger)

	// Setup router
	router := router.Init(usecases, logger)

	// Start server
	httpServer := httpv1.Init(cfg.HTTPServerConfig, router.Handler())

	<-make(chan os.Signal, 1)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	httpServer.Shutdown(ctx)

}
