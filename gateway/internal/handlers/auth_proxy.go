package handlers

import (
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewAuthProxy(targetURL string, logger *slog.Logger) http.Handler {
	target, err := url.Parse(targetURL)
	if err != nil {
		logger.Error("Failed to parse auth service URL", slog.String("url", targetURL), slog.Any("error", err))
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		logger.Error("Auth service proxy error",
			slog.String("path", r.URL.Path),
			slog.Any("error", err),
		)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(`{"error":"auth service unavailable"}`))
	}

	return proxy
}
