package httpserver

import (
	"auth/config"
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func Init(cfg config.HTTPServerConfig, router http.Handler, logger *slog.Logger) *Server {
	httpServer := &http.Server{
		Addr:    cfg.Address + ":" + cfg.Port,
		Handler: router,
	}

	addr := fmt.Sprintf("%s:%s", cfg.Address, cfg.Port)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server start error: %s\n", err)
		}
	}()

	slog.Info("Server started", slog.String("Address", addr))

	return &Server{
		httpServer: httpServer,
	}
}

func (server *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.httpServer.Shutdown(ctx)
	slog.Info("Server shuted down")
}
