package usecase

import (
	"auth/internal/application/dto"
	"auth/internal/domain/user/entity"
	"auth/internal/domain/user/repository"
	"log/slog"
)

type GetUserByUUIDUseCase struct {
	postgres repository.Postgres
	logger   *slog.Logger
}

func NewGetUserByUUIDUseCase(postgres repository.Postgres, logger *slog.Logger) *GetUserByUUIDUseCase {
	return &GetUserByUUIDUseCase{
		postgres: postgres,
		logger:   logger,
	}
}

func (uc *GetUserByUUIDUseCase) Execute(uuid string) (*dto.UserOutput, error) {
	const op = "application.usecase.GetUserByUUIDUseCase.Execute"
	user, err := uc.postgres.GetByUUID(uuid)

	if err != nil {
		if err == entity.UserNotFoundError {
			uc.logger.Warn("User not found", slog.String("uuid", uuid))
			return nil, err
		}
		uc.logger.Error("Failed to get user", slog.String("op", op))
		return nil, err
	}

	return &dto.UserOutput{
		UUID:      user.UUID.String(),
		Email:     user.Email.String(),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
