package auth

import (
	"auth/internal/domain/entity"
	"auth/internal/domain/interfaces"
	"auth/internal/usecase/auth/dto"
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
	token, err := uc.authRepo.ValidateToken(input.Refresh)
	if err != nil {
		return nil, entity.ErrInvalidToken
	}

	isTokenExisting, err := uc.authRepo.CheckTokenExisting(token.TokenID)
	if err != nil || !isTokenExisting {
		return nil, entity.ErrTokenNotFound
	}

	newAccess, err := uc.authRepo.GenerateToken(entity.AccessTokenType, token.UserUUID)
	if err != nil {
		uc.logger.Error("Failed to generate access token", "error", err)
		return nil, err
	}

	newRefresh, err := uc.authRepo.GenerateToken(entity.RefreshTokenType, token.UserUUID)
	if err != nil {
		uc.logger.Error("Failed to generate refresh token", "error", err)
		return nil, err
	}

	err = uc.authRepo.SaveToken(newRefresh)
	if err != nil {
		uc.logger.Error("Failed to save new refresh token", "error", err)
		return nil, err
	}

	err = uc.authRepo.RevokeToken(token.TokenID)
	if err != nil {
		uc.logger.Error("Failed to revoke old refresh token", "error", err)
		return nil, err
	}

	return &dto.AuthOutput{
		Access:  newAccess.Token,
		Refresh: newRefresh.Token,
	}, nil

}
