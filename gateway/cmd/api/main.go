package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"gateway/config"
	router "gateway/internal"
	"gateway/pkg/httpserver"
	"gateway/pkg/logger"
)

func main() {
	cfg := config.Init()
	log := logger.Init(cfg.Env)

	AppRun(cfg, log)
}

func AppRun(cfg config.Config, log *slog.Logger) {
	ginRouter := router.NewRouter(cfg, log)

	httpServer := httpserver.Init(cfg.HTTPServerConfig, ginRouter, log)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	<-quit

	httpServer.Shutdown()

	os.Exit(0)
}
