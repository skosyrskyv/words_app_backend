package usecase

import (
	"auth/internal/application/dto"
	"auth/internal/domain/user/entity"
	"auth/internal/domain/user/repository"
	"log/slog"
)

type CreateUserUseCase struct {
	postgres repository.Postgres
	logger   *slog.Logger
}

func NewCreateUserUseCase(
	postgres repository.Postgres,
	logger *slog.Logger,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		postgres: postgres,
		logger:   logger,
	}
}

func (uc *CreateUserUseCase) Execute(input dto.CrateUserInput) (*dto.UserOutput, error) {

	foundUser, err := uc.postgres.GetByEmail(input.Email)

	if err == nil && foundUser != nil {
		uc.logger.Warn("User already exist", slog.String("email", input.Email))
		return nil, entity.EmailAlreadyExistsError
	}

	user, err := entity.NewUser(input.Email, input.Password)

	if err != nil {
		uc.logger.Error("Failed to create user")
		return nil, err
	}

	createdUser, err := uc.postgres.Create(user)

	if err != nil {
		uc.logger.Error("Failed to save user to DB")
		return nil, err
	}

	return &dto.UserOutput{
		UUID:      createdUser.UUID.String(),
		Email:     createdUser.Email.String(),
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}, nil
}
