package usecase

import (
	"auth/internal/domain/interfaces"
	"auth/internal/usecase/auth"
	"auth/internal/usecase/user"
	"log/slog"
)

type UseCases struct {
	User *UserUseCases
	Auth *AuthUseCases
}

type UserUseCases struct {
	CreateUser    *user.CreateUserUseCase
	GetUserByUUID *user.GetUserByUUIDUseCase
}

type AuthUseCases struct {
	Login   *auth.LoginUseCase
	Refresh *auth.RefreshTokenUseCase
}

func NewUseCases(user *UserUseCases, auth *AuthUseCases) *UseCases {
	return &UseCases{
		User: user,
		Auth: auth,
	}
}

func NewUserUseCases(
	repo interfaces.UserRepository,
	logger *slog.Logger,
) *UserUseCases {
	return &UserUseCases{
		CreateUser:    user.NewCreateUserUseCase(repo, logger),
		GetUserByUUID: user.NewGetUserByUUIDUseCase(repo, logger),
	}
}

func NewAuthUseCases(
	authRepo interfaces.AuthRepository,
	userRepo interfaces.UserRepository,
	logger *slog.Logger,
) *AuthUseCases {
	return &AuthUseCases{
		Login:   auth.NewLoginUseCase(authRepo, userRepo, logger),
		Refresh: auth.NewRefreshTokenUsecase(authRepo, userRepo, logger),
	}
}
