package handlers

import (
	"auth/internal/domain/entity"
	httpDto "auth/internal/presentation/http/dto"
	"auth/internal/usecase"
	"auth/internal/usecase/user/dto"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
			Error:   http.StatusText(http.StatusBadRequest),
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
				Error:   http.StatusText(http.StatusConflict),
				Message: "User with this email already exists",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, httpDto.ErrorResponse{
			Error:   http.StatusText(http.StatusInternalServerError),
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

func (h *UserHandler) GetUserMe(c *gin.Context) {
	uuidParam := c.GetHeader("X-User-UUID")

	if _, err := uuid.Parse(uuidParam); err != nil {
		c.JSON(http.StatusBadRequest, httpDto.ErrorResponse{
			Error:   http.StatusText(http.StatusBadRequest),
			Message: "Invalid UUID format",
		})
		return
	}

	userOutput, err := h.useCases.GetUserByUUID.Execute(uuidParam)

	if err != nil {
		if err == entity.ErrUserNotFound {
			c.JSON(http.StatusNotFound, httpDto.ErrorResponse{
				Error:   http.StatusText(http.StatusNotFound),
				Message: "User with this UUID not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, httpDto.ErrorResponse{
			Error:   http.StatusText(http.StatusInternalServerError),
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

func (h *UserHandler) GetUserByUUID(c *gin.Context) {
	uuidParam := c.Param("uuid")

	if _, err := uuid.Parse(uuidParam); err != nil {
		c.JSON(http.StatusBadRequest, httpDto.ErrorResponse{
			Error:   http.StatusText(http.StatusBadRequest),
			Message: "Invalid UUID format",
		})
		return
	}

	userOutput, err := h.useCases.GetUserByUUID.Execute(uuidParam)

	if err != nil {
		if err == entity.ErrUserNotFound {
			c.JSON(http.StatusNotFound, httpDto.ErrorResponse{
				Error:   http.StatusText(http.StatusNotFound),
				Message: "User with this UUID not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, httpDto.ErrorResponse{
			Error:   http.StatusText(http.StatusInternalServerError),
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
