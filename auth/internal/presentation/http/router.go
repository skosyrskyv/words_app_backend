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

	// User routes
	userGroup := router.Group("/api/v1/users")
	{
		userGroup.GET("/:uuid", userHandler.GetUserByUUID)
		userGroup.POST("/registration", userHandler.CreateUser)
	}

	authGroup := router.Group("/api/v1/auth")
	{
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/token/refresh", authHandler.RefreshToken)
	}

	return router
}
