package handlers

import (
	"auth/internal/domain/entity"
	httpDto "auth/internal/presentation/http/dto"
	"auth/internal/usecase"
	"auth/internal/usecase/user/dto"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	useCases *usecase.UserUseCases
	logger   *slog.Logger
}

func NewUserHandler(useCases *usecase.UserUseCases, logger *slog.Logger) *UserHandler {
	return &UserHandler{
		useCases: useCases,
		logger:   logger,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	req := httpDto.CreateUserRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, httpDto.ErrorResponse{
			Error:   "INVALID_REQUEST",
			Message: err.Error(),
		})
		return
	}

	input := dto.CreateUserInput{
		Email:    req.Email,
		Password: req.Password,
	}

	userOutput, err := h.useCases.CreateUser.Execute(input)

	if err != nil {
		if err == entity.ErrEmailAlreadyExists {
			c.JSON(http.StatusConflict, httpDto.ErrorResponse{
				Error:   "EMAIL_ALREADY_EXISTS",
				Message: "User with this email already exists",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, httpDto.ErrorResponse{
			Error:   "INTERNAL_SERVER_ERROR",
			Message: "Failed to create user",
		})
		return
	}

	response := httpDto.UserResponse{
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
		c.JSON(http.StatusBadRequest, httpDto.ErrorResponse{
			Error:   "INVALID_REQUEST",
			Message: "UUID is required",
		})
		return
	}

	userOutput, err := h.useCases.GetUserByUUID.Execute(uuid)

	if err != nil {
		if err == entity.ErrUserNotFound {
			c.JSON(http.StatusNotFound, httpDto.ErrorResponse{
				Error:   "USER_NOT_FOUND",
				Message: "User with this UUID not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, httpDto.ErrorResponse{
			Error:   "INTERNAL_SERVER_ERROR",
			Message: "Failed to get user",
		})
		return
	}

	response := httpDto.UserResponse{
		UUID:      userOutput.UUID,
		Email:     userOutput.Email,
		CreatedAt: userOutput.CreatedAt,
		UpdatedAt: userOutput.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}
