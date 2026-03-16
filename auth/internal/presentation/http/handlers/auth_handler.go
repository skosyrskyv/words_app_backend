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

// Login godoc
// @Summary      Login
// @Description  Authenticate user and receive access + refresh tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginRequest true "Login credentials"
// @Success      200 {object} dto.AuthSuccessResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /api/v1/auth/login [post]
func (h *handler) Login(c *gin.Context) {
	req := httpdto.LoginRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, httpdto.ErrorResponse{
			Error:   http.StatusText(http.StatusBadRequest),
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
				Error:   http.StatusText(http.StatusNotFound),
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, httpdto.ErrorResponse{
			Error:   http.StatusText(http.StatusInternalServerError),
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

// RefreshToken godoc
// @Summary      Refresh token
// @Description  Get a new access token using a refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body dto.RefreshTokenRequest true "Refresh token"
// @Success      200 {object} dto.AuthSuccessResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      401 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /api/v1/auth/token/refresh [post]
func (h *handler) RefreshToken(c *gin.Context) {
	req := httpdto.RefreshTokenRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, httpdto.ErrorResponse{
			Error:   http.StatusText(http.StatusBadRequest),
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
				Error:   http.StatusText(http.StatusUnauthorized),
				Message: "The provided token is invalid.",
			})
			return
		}
		if err == entity.ErrTokenNotFound {
			c.JSON(http.StatusNotFound, httpdto.ErrorResponse{
				Error:   http.StatusText(http.StatusNotFound),
				Message: "The provided token was not found.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, httpdto.ErrorResponse{
			Error:   http.StatusText(http.StatusInternalServerError),
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
