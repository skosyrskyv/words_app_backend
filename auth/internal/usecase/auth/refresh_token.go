package auth

import (
	"auth/internal/domain/interfaces"
	"auth/internal/usecase/auth/dto"
	"log/slog"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshTokenUseCase struct {
	authRepo interfaces.AuthRepository
	userRepo interfaces.UserRepository
	logger   *slog.Logger
}

func NewRefreshTokenUsecase(authRepo interfaces.AuthRepository, userRepo interfaces.UserRepository, logger *slog.Logger) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{
		authRepo: authRepo,
		userRepo: userRepo,
		logger:   logger,
	}
}

func (uc *RefreshTokenUseCase) Execute(input dto.RefreshTokenInput) (*dto.AuthOutput, error) {
	access := jwt.New(jwt.SigningMethodES256)
	refresh := jwt.New(jwt.SigningMethodES256)

	return &dto.AuthOutput{
		Access:  access.Raw,
		Refresh: refresh.Raw,
	}, nil
}
