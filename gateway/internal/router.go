package internal

import (
	"gateway/config"
	"gateway/internal/handlers"
	"gateway/internal/middleware"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(cfg config.Config, logger *slog.Logger) *gin.Engine {
	router := gin.Default()

	authProxy := handlers.NewAuthProxy(cfg.ServicesConfig.AuthServiceURL, logger)
	jwtMiddleware := middleware.NewJWTMiddleware(cfg.JWTConfig, logger)

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Static placeholder
	router.StaticFile("/", "/app/static/index.html")

	apiGroup := router.Group("/api")

	// Version
	versionGroup := apiGroup.Group("/v1")

	// Public auth routes (login, registration, token refresh — без токена)
	versionGroup.Any("/auth/*path", gin.WrapH(authProxy))

	// Protected routes (требуют валидный access token)
	protected := versionGroup.Group("/users")
	protected.Use(jwtMiddleware.RequireAuth())
	{
		protected.Any("/*path", gin.WrapH(authProxy))
	}

	return router
}
