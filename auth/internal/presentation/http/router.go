package router

import (
	"auth/internal/presentation/http/handlers"
	"auth/internal/usecase"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func NewRouter(useCases *usecase.UseCases, logger *slog.Logger) *gin.Engine {
	router := gin.Default()

	userHandler := handlers.NewUserHandler(useCases.User, logger)
	authHandler := handlers.NewAuthHandler(useCases.Auth, logger)

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API
	apiGroup := router.Group("/api")

	// Version
	versionGroup := apiGroup.Group("/v1")

	// User routes
	userGroup := versionGroup.Group("/users")
	{
		userGroup.GET("/:uuid", userHandler.GetUserByUUID)
		userGroup.GET("/me", userHandler.GetUserMe)
	}

	// Auth routes
	authGroup := versionGroup.Group("/auth")
	{
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/registration", userHandler.CreateUser)
		authGroup.POST("/token/refresh", authHandler.RefreshToken)
	}

	return router
}
