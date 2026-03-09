package httpv1

import (
	"auth/config"
	"log"
	"net/http"
)

func Init(cfg config.HTTPServerConfig, router http.Handler) *http.Server {
	httpServer := &http.Server{
		Addr:    cfg.Port,
		Handler: router,
	}

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server start error: %s\n", err)
	}

	return httpServer
}
