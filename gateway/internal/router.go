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

	// CORS configuration
	router.Use(cfg.Cors)

	authProxy := handlers.NewAuthProxy(cfg.ServicesConfig.AuthServiceURL, logger)
	translations := handlers.NewTranslationsHandler(cfg.ServicesConfig.TranslationsServiceAddr, logger)
	jwtMiddleware := middleware.NewJWTMiddleware(cfg.JWTConfig, logger)

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Static placeholder
	router.StaticFile("/", "/app/static/index.html")

	apiGroup := router.Group("/api")

	// Version
	v1Group := apiGroup.Group("/v1")
	{
		v1Group.StaticFile("/swagger", "/app/static/swagger.html")
		v1Group.GET("/swagger/*any", gin.WrapH(authProxy))
	}

	// Public routes
	public := v1Group.Group("/")
	{
		public.Any("/auth/*any", gin.WrapH(authProxy))
	}

	// Optional auth routes
	optional := v1Group.Group("/")
	optional.Use(jwtMiddleware.OptionalAuth())
	{
		optional.POST("/translations/translate", translations.Translate)
	}

	// Protected routes
	protected := v1Group.Group("/")
	protected.Use(jwtMiddleware.RequireAuth())
	{
		protected.Any("/users/*any", gin.WrapH(authProxy))
	}

	return router
}
