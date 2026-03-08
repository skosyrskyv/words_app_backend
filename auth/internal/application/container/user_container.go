package container

import (
	"log/slog"

	usecase "auth/internal/application/usecase/user"
	"auth/internal/domain/ports/repository"
)

type UserUseCases struct {
	CreateUser    *usecase.CreateUserUseCase
	GetUserByUUID *usecase.GetUserByUUIDUseCase
	// ChangePassword *usecase.ChangePasswordUseCase
	// DeleteUser     *usecase.DeleteUserUseCase
}

func NewUserUseCases(
	userRepository repository.UserRepository,
	logger *slog.Logger,
) *UserUseCases {
	return &UserUseCases{
		CreateUser:    usecase.NewCreateUserUseCase(userRepository, logger),
		GetUserByUUID: usecase.NewGetUserByUUIDUseCase(userRepository, logger),
		// ChangePassword: usecase.NewChangePasswordUseCase(userRepository, logger),
		// DeleteUser:     usecase.NewDeleteUserUseCase(userRepository, logger),
	}
}
