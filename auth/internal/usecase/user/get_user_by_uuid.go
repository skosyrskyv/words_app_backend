package user

import (
	"auth/internal/domain/entity"
	"auth/internal/domain/interfaces"
	"auth/internal/usecase/user/dto"
	"log/slog"
)

type GetUserByUUIDUseCase struct {
	repo   interfaces.UserRepository
	logger *slog.Logger
}

func NewGetUserByUUIDUseCase(repo interfaces.UserRepository, logger *slog.Logger) *GetUserByUUIDUseCase {
	return &GetUserByUUIDUseCase{
		repo:   repo,
		logger: logger,
	}
}

func (uc *GetUserByUUIDUseCase) Execute(uuid string) (*dto.UserOutput, error) {
	user, err := uc.repo.GetUserByUUID(uuid)

	if err != nil {
		if err == entity.UserNotFoundError {
			uc.logger.Warn("User not found", slog.String("uuid", uuid))
			return nil, err
		}
		uc.logger.Error("Failed to get user")
		return nil, err
	}

	if user.DeletedAt.Valid {
		uc.logger.Warn("User is deleted", slog.String("uuid", uuid), slog.Time("deleted_at", user.DeletedAt.Time))
		return nil, entity.UserNotFoundError
	}

	return &dto.UserOutput{
		UUID:      user.UUID.String(),
		Email:     user.Email.String(),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
