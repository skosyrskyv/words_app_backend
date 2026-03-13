package user

import (
	"auth/internal/domain/entity"
	"auth/internal/domain/interfaces"
	"auth/internal/usecase/user/dto"
	"log/slog"
)

type CreateUserUseCase struct {
	repo   interfaces.UserRepository
	logger *slog.Logger
}

func NewCreateUserUseCase(
	repo interfaces.UserRepository,
	logger *slog.Logger,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo:   repo,
		logger: logger,
	}
}

func (uc *CreateUserUseCase) Execute(input dto.CreateUserInput) (*dto.UserOutput, error) {

	foundUser, err := uc.repo.GetUserByEmail(input.Email)

	if err == nil && foundUser != nil {
		uc.logger.Warn("User already exist", slog.String("email", input.Email))
		return nil, entity.ErrEmailAlreadyExists
	}

	user, err := entity.NewUser(input.Email, input.Password)

	if err != nil {
		uc.logger.Error("Failed to create user")
		return nil, err
	}

	createdUser, err := uc.repo.CreateUser(user)

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
