package handlers

import (
	"auth/internal/domain/entity"
	httpdto "auth/internal/presentation/http/dto"
	"auth/internal/usecase/auth/dto"

	"auth/internal/usecase"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	useCases *usecase.AuthUseCases
	logger   *slog.Logger
}

func NewAuthHandler(useCases *usecase.AuthUseCases, logger *slog.Logger) *handler {
	return &handler{
		useCases: useCases,
		logger:   logger,
	}
}

func (h *handler) Login(c *gin.Context) {
	req := httpdto.LoginRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, httpdto.ErrorResponse{
			Error:   "INVALID_REQUEST",
			Message: err.Error(),
		})
		return
	}

	input := dto.AuthInput{
		Email:    req.Email,
		Password: req.Password,
	}

	output, err := h.useCases.Login.Execute(input)

	if err != nil {
		if err == entity.ErrUserCredentials {
			c.JSON(http.StatusNotFound, httpdto.ErrorResponse{
				Error:   "NO_USER_CREDENTIALS",
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, httpdto.ErrorResponse{
			Error:   "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
		})
		return
	}

	response := httpdto.AuthSuccessResponse{
		Access:  output.Access,
		Refresh: output.Refresh,
	}

	c.JSON(http.StatusOK, response)
}

func (h *handler) RefreshToken(c *gin.Context) {
	req := httpdto.RefreshTokenRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, httpdto.ErrorResponse{
			Error:   "INVALID_REQUEST",
			Message: err.Error(),
		})
		return
	}

	input := dto.RefreshTokenInput{
		Refresh: req.Refresh,
	}

	output, err := h.useCases.Refresh.Execute(input)
	if err != nil {
		if err == entity.ErrInvalidToken {
			c.JSON(http.StatusUnauthorized, httpdto.ErrorResponse{
				Error:   "INVALID_TOKEN",
				Message: "The provided token is invalid.",
			})
			return
		}
		if err == entity.ErrTokenNotFound {
			c.JSON(http.StatusNotFound, httpdto.ErrorResponse{
				Error:   "TOKEN_NOT_FOUND",
				Message: "The provided token was not found.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, httpdto.ErrorResponse{
			Error:   "INTERNAL_SERVER_ERROR",
			Message: "An internal server error occurred.",
		})
		return
	}

	response := httpdto.AuthSuccessResponse{
		Access:  output.Access,
		Refresh: output.Refresh,
	}

	c.JSON(http.StatusOK, response)
}
