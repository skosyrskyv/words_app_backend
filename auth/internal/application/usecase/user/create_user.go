package usecase

import (
	"auth/internal/application/dto"
	"auth/internal/domain/entity"
	"auth/internal/domain/ports/repository"
	"log/slog"
)

type CreateUserUseCase struct {
	userRepository repository.UserRepository
	logger         *slog.Logger
}

func NewCreateUserUseCase(
	userRepository repository.UserRepository,
	logger *slog.Logger,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (uc *CreateUserUseCase) Execute(input dto.CrateUserInput) (*dto.UserOutput, error) {
	const op = "application.usecase.CreateUserUseCase.Execute"

	foundUser, err := uc.userRepository.GetByEmail(input.Email)

	if err != nil && foundUser != nil {
		uc.logger.Warn("User already exist", slog.String("email", input.Email))
		return nil, entity.EmailAlreadyExistsError
	}

	user, err := entity.NewUser(input.Email, input.Password)

	if err != nil {
		uc.logger.Error("Failed to create user", slog.String("op", op))
		return nil, err
	}

	userUUID, err := uc.userRepository.Create(user)

	if err != nil {
		uc.logger.Error("Failed to save user to DB", slog.String("op", op))
		return nil, err
	}

	return &dto.UserOutput{
		UUID:  userUUID,
		Email: input.Email,
	}, nil
}
