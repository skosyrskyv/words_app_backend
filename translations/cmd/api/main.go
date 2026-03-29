package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"translations/config"
	"translations/internal/infrastructure/translation"
	grpcrouter "translations/internal/presentation/grpc"
	usecases "translations/internal/usecase"
	"translations/pkg/grpcserver"
	"translations/pkg/httpclient"
	"translations/pkg/logger"
	"translations/pkg/postgres"
)

func main() {
	cfg := config.Init()
	logger := logger.Init(cfg.Env)

	AppRun(context.Background(), cfg, logger)
}

func AppRun(ctx context.Context, cfg config.Config, logger *slog.Logger) {
	// // Initialize Redis
	// redis := redis.Init(ctx, cfg.RedisConfig)

	// Initialize Postgres
	postgres := postgres.Init(cfg.PostgresConfig)

	// Http Client
	http := httpclient.NewHTTPClient()

	// Initialize repositories
	translationsRepository := translation.NewRepository(http, postgres)

	// Initialize UseCases
	translationsUseCase := usecases.NewTranslationUseCases(translationsRepository, logger)
	usecases := usecases.NewUseCases(translationsUseCase)

	// Initialize gRPC server
	router := grpcrouter.NewGRPCRouter(usecases)
	grpcServer := grpcserver.Init(cfg.GRPCServerConfig, router)

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	<-quit

	grpcServer.Shutdown()
	postgres.Close()
	os.Exit(0)
}
