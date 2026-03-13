package logger

import (
	"auth/config"
	"log"
	"log/slog"
	"os"
)

func Init(env string) *slog.Logger {
	var handler slog.Handler

	switch env {
	case config.LOCAL:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true})
	case config.DEV:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case config.PROD:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	default:
		log.Fatal("Incorrect ENV type")
	}

	return slog.New(handler)
}
