package logger

import (
	envtypes "auth/internal/const/envType"
	"log/slog"
	"os"
)

func Setup(env string) {
	var handler slog.Handler
	var logger *slog.Logger

	switch env {
	case envtypes.LOCAL:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true})
	case envtypes.DEV:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case envtypes.PROD:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}

	logger = slog.New(handler)
	slog.SetDefault(logger)
}
