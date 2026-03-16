// @title           Words App API
// @version         1.0
// @description     REST API for Words App — accessible through the API Gateway

// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description                Format: Bearer <access_token>

package main

import (
	"auth/config"
	"auth/internal/infrastructure/auth"
	"auth/internal/infrastructure/user"
	router "auth/internal/presentation/http"
	"auth/internal/usecase"
	"os/signal"
	"syscall"

	"auth/pkg/httpserver"
	"auth/pkg/logger"
	"auth/pkg/postgres"
	"auth/pkg/redis"
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
	// Initialize Redis
	redis := redis.Init(ctx, cfg.RedisConfig)

	// Initialize Postgres
	postgres := postgres.Init(cfg.PostgresConfig)

	// Initialize repositories
	userRepository := user.NewRepository(postgres.DB())
	authRepository := auth.NewRepository(cfg.JWTConfig, redis, postgres)

	// Initialize UseCases
	userUC := usecase.NewUserUseCases(userRepository, logger)
	authUC := usecase.NewAuthUseCases(authRepository, userRepository, logger)
	useCases := usecase.NewUseCases(userUC, authUC)

	// Setup router
	ginRouter := router.NewRouter(useCases, logger)

	// Start server
	httpServer := httpserver.Init(cfg.HTTPServerConfig, ginRouter, logger)

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	<-quit

	httpServer.Shutdown()
	postgres.Close()
	redis.Close()

	os.Exit(0)
}
