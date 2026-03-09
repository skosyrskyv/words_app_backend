package handlers

import (
	"auth/internal/application/dto"
	"auth/internal/application/usecase"
	"auth/internal/domain/user/entity"
	httpdto "auth/internal/presentation/http/dto"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecases *usecase.UserUseCases
	logger   *slog.Logger
}

func NewUserHandler(usecases *usecase.UserUseCases, logger *slog.Logger) *UserHandler {
	return &UserHandler{
		usecases: usecases,
		logger:   logger,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	req := httpdto.CreateUserRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Invalid request", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, httpdto.ErrorResponse{
			Error:   "INVALID_REQUEST",
			Message: err.Error(),
		})
		return
	}

	input := dto.CrateUserInput{
		Email:    req.Email,
		Password: req.Password,
	}

	userOutput, err := h.usecases.CreateUser.Execute(input)

	if err != nil {
		h.logger.Error("Failed to create user", slog.String("error", err.Error()))

		if err == entity.EmailAlreadyExistsError {
			c.JSON(http.StatusConflict, httpdto.ErrorResponse{
				Error:   "EMAIL_ALREADY_EXISTS",
				Message: "User with this email already exists",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, httpdto.ErrorResponse{
			Error:   "INTERNAL_ERROR",
			Message: "Failed to create user",
		})
		return
	}

	response := httpdto.UserResponse{
		UUID:      userOutput.UUID,
		Email:     userOutput.Email,
		CreatedAt: userOutput.CreatedAt,
		UpdatedAt: userOutput.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) GetUserByUUID(c *gin.Context) {
	uuid := c.Param("uuid")

	if uuid == "" {
		h.logger.Warn("UUID not provided")
		c.JSON(http.StatusBadRequest, httpdto.ErrorResponse{
			Error:   "INVALID_REQUEST",
			Message: "UUID is required",
		})
		return
	}

	userOutput, err := h.usecases.GetUserByUUID.Execute(uuid)

	if err != nil {
		h.logger.Warn("Failed to get user", slog.String("uuid", uuid), slog.String("error", err.Error()))

		if err == entity.UserNotFoundError {
			c.JSON(http.StatusNotFound, httpdto.ErrorResponse{
				Error:   "USER_NOT_FOUND",
				Message: "User with this UUID not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, httpdto.ErrorResponse{
			Error:   "INTERNAL_ERROR",
			Message: "Failed to get user",
		})
		return
	}

	response := httpdto.UserResponse{
		UUID:      userOutput.UUID,
		Email:     userOutput.Email,
		CreatedAt: userOutput.CreatedAt,
		UpdatedAt: userOutput.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}
