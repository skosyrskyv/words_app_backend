package repositories

import (
	"auth/internal/domain/user/entity"
	"auth/internal/domain/user/entity/vo"
	"auth/internal/infrastructure/postgres/models"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *entity.User) (*entity.User, error) {
	model := &models.User{
		Uuid:     user.UUID.String(),
		Email:    user.Email.String(),
		Password: user.PasswordHash.String(),
	}

	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}

	return modelToEntity(model)
}

func (r *UserRepository) GetByUUID(uuid string) (*entity.User, error) {
	var model models.User

	if err := r.db.Where("uuid = ?", uuid).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.UserNotFoundError
		}
		return nil, err
	}

	return modelToEntity(&model)
}

func (r *UserRepository) GetByEmail(email string) (*entity.User, error) {
	var model models.User

	if err := r.db.Where("email = ?", email).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return modelToEntity(&model)
}

func (r *UserRepository) UpdatePassword(user *entity.User) error {
	model := &models.User{
		Password: user.PasswordHash.String(),
	}

	return r.db.Model(&models.User{}).Where("uuid = ?", user.UUID.String()).Update("password", model.Password).Error
}

func (r *UserRepository) Delete(uuid string) error {
	return r.db.Delete(&models.User{}, "uuid = ?", uuid).Error
}

func modelToEntity(model *models.User) (*entity.User, error) {
	uid, err := uuid.Parse(model.Uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to parse UUID: %w", err)
	}

	email, err := vo.NewEmailFromString(model.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to create email value object: %w", err)
	}

	passwordHash, err := vo.NewPasswordHashFromString(model.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to create password hash value object: %w", err)
	}

	return &entity.User{
		UUID:         uid,
		Email:        email,
		PasswordHash: passwordHash,
		IsActive:     model.IsActive,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		DeletedAt:    model.DeletedAt,
	}, nil
}
