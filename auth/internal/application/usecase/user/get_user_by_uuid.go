package usecase

import (
	"auth/internal/application/dto"
	"auth/internal/domain/entity"
	"auth/internal/domain/ports/repository"
	"log/slog"
)

type GetUserByUUIDUseCase struct {
	repository repository.UserRepository
	logger     *slog.Logger
}

func NewGetUserByUUIDUseCase(repository repository.UserRepository, logger *slog.Logger) *GetUserByUUIDUseCase {
	return &GetUserByUUIDUseCase{
		repository: repository,
		logger:     logger,
	}
}

func (uc *GetUserByUUIDUseCase) Execute(uuid string) (*dto.UserOutput, error) {
	const op = "application.usecase.GetUserByUUIDUseCase.Execute"
	user, err := uc.repository.GetByUUID(uuid)

	if err != nil {
		if err == entity.UserNotFoundError {
			uc.logger.Warn("User not found", slog.String("uuid", uuid))
			return nil, err
		}
		uc.logger.Error("Failed to get user", slog.String("op", op))
		return nil, err
	}

	return &dto.UserOutput{
		UUID:  user.UUID.String(),
		Email: user.Email.String(),
	}, nil
}
