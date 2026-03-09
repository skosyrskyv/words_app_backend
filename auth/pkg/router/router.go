package router

import (
	"auth/internal/application/usecase"
	"auth/internal/presentation/http/handlers"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func Init(usecases *usecase.UserUseCases, logger *slog.Logger) *gin.Engine {
	router := gin.Default()

	userHandler := handlers.NewUserHandler(usecases, logger)

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

	return router
}
