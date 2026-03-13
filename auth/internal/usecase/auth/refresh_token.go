package auth

import (
	"auth/internal/domain/interfaces"
	"auth/internal/usecase/auth/dto"
	"errors"
	"log/slog"
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
	return nil, errors.ErrUnsupported
}
