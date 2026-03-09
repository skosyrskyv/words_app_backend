package main

import (
	"auth/config"
	"auth/internal/application/usecase"
	"auth/internal/infrastructure/postgres/repositories"
	"auth/pkg/httpserver"
	"auth/pkg/logger"
	"auth/pkg/postgres"
	"auth/pkg/router"
	"context"
	"log/slog"
	"os"
)

func main() {

	cfg := config.Init()

	logger := logger.Init(cfg.Env)

	AppRun(context.Background(), cfg, logger)
}

func AppRun(ctx context.Context, cfg config.Config, logger *slog.Logger) {

	// Initialize postgres
	postgres := postgres.Init(cfg.PostgresConfig)

	// Initialize repositories
	userRepository := repositories.NewUserRepository(postgres.DB())

	// Initialize usecases
	usecases := usecase.NewUserUseCases(userRepository, logger)

	// Setup router
	ginRouter := router.Init(usecases, logger)

	// Start server
	httpServer := httpserver.Init(cfg.HTTPServerConfig, ginRouter, logger)

	<-make(chan os.Signal, 1)

	postgres.Close()
	httpServer.Close()
}
