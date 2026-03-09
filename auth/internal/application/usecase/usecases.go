package usecase

import (
	"auth/internal/domain/user/repository"
	"log/slog"
)

type UserUseCases struct {
	CreateUser    *CreateUserUseCase
	GetUserByUUID *GetUserByUUIDUseCase
	// ChangePassword *usecase.ChangePasswordUseCase
	// DeleteUser     *usecase.DeleteUserUseCase
}

func NewUserUseCases(
	postgres repository.Postgres,
	logger *slog.Logger,
) *UserUseCases {
	return &UserUseCases{
		CreateUser:    NewCreateUserUseCase(postgres, logger),
		GetUserByUUID: NewGetUserByUUIDUseCase(postgres, logger),
		// ChangePassword: usecase.NewChangePasswordUseCase(userRepository, logger),
		// DeleteUser:     usecase.NewDeleteUserUseCase(userRepository, logger),
	}
}
