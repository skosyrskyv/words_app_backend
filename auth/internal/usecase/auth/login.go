package auth

import (
	"auth/internal/domain/entity"
	"auth/internal/domain/interfaces"
	"auth/internal/usecase/auth/dto"

	"log/slog"
)

type LoginUseCase struct {
	authRepo interfaces.AuthRepository
	userRepo interfaces.UserRepository
	logger   *slog.Logger
}

func NewLoginUseCase(authRepo interfaces.AuthRepository, userRepo interfaces.UserRepository, logger *slog.Logger) *LoginUseCase {
	return &LoginUseCase{
		authRepo: authRepo,
		userRepo: userRepo,
		logger:   logger,
	}
}

func (uc *LoginUseCase) Execute(input dto.AuthInput) (*dto.AuthOutput, error) {
	// Get users
	user, err := uc.userRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, entity.ErrUserCredentials
	}

	// Check passwords
	res, err := user.VerifyPassword(input.Password)
	if err != nil || !res {
		return nil, entity.ErrUserCredentials
	}

	// Generate tokens
	access, err := uc.authRepo.GenerateToken(entity.AccessTokenType, user.UUID.String())
	if err != nil {
		uc.logger.Error("Failed to generate access token", "error", err)
		return nil, err
	}

	refresh, err := uc.authRepo.GenerateToken(entity.RefreshTokenType, user.UUID.String())
	if err != nil {
		uc.logger.Error("Failed to generate refresh token", "error", err)
		return nil, err
	}

	// Save refresh token
	err = uc.authRepo.SaveToken(refresh)
	if err != nil {
		uc.logger.Error("Failed to save refresh token", "error", err)
		return nil, err
	}

	return &dto.AuthOutput{
		Access:  access.Token,
		Refresh: refresh.Token,
	}, nil
}
