package user

import (
	"auth/internal/domain/entity"
	"auth/internal/infrastructure/user/model"

	"gorm.io/gorm"
)

type repository struct {
	postgres *gorm.DB
}

func NewRepository(postgres *gorm.DB) *repository {
	return &repository{postgres: postgres}
}

func (r *repository) CreateUser(user *entity.User) (*entity.User, error) {
	model := &model.UserModel{
		Uuid:     user.UUID.String(),
		Email:    user.Email.String(),
		Password: user.PasswordHash.String(),
	}

	if err := r.postgres.Create(model).Error; err != nil {
		return nil, err
	}

	return model.ToEntity()
}

func (r *repository) GetUserByUUID(uuid string) (*entity.User, error) {
	var model model.UserModel

	if err := r.postgres.Where("uuid = ?", uuid).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.ErrUserNotFound
		}
		return nil, err
	}

	return model.ToEntity()
}

func (r *repository) GetUserByEmail(email string) (*entity.User, error) {
	var model model.UserModel

	if err := r.postgres.Where("email = ?", email).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.ErrUserCredentials
		}
		return nil, err
	}

	return model.ToEntity()
}

func (r *repository) UpdateUserPassword(user *entity.User) error {
	m := &model.UserModel{
		Password: user.PasswordHash.String(),
	}

	return r.postgres.Model(&model.UserModel{}).Where("uuid = ?", user.UUID.String()).Update("password", m.Password).Error
}

func (r *repository) DeleteUser(uuid string) error {
	return r.postgres.Delete(&model.UserModel{}, "uuid = ?", uuid).Error
}
