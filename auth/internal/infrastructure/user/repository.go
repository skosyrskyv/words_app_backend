package user

import (
	"auth/internal/domain/entity"
	"auth/internal/infrastructure/user/model"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(user *entity.User) (*entity.User, error) {
	model := &model.UserModel{
		Uuid:     user.UUID.String(),
		Email:    user.Email.String(),
		Password: user.PasswordHash.String(),
	}

	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}

	return model.ToEntity()
}

func (r *repository) GetUserByUUID(uuid string) (*entity.User, error) {
	var model model.UserModel

	if err := r.db.Where("uuid = ?", uuid).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.UserNotFoundError
		}
		return nil, err
	}

	return model.ToEntity()
}

func (r *repository) GetUserByEmail(email string) (*entity.User, error) {
	var model model.UserModel

	if err := r.db.Where("email = ?", email).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return model.ToEntity()
}

func (r *repository) UpdateUserPassword(user *entity.User) error {
	m := &model.UserModel{
		Password: user.PasswordHash.String(),
	}

	return r.db.Model(&model.UserModel{}).Where("uuid = ?", user.UUID.String()).Update("password", m.Password).Error
}

func (r *repository) DeleteUser(uuid string) error {
	return r.db.Delete(&model.UserModel{}, "uuid = ?", uuid).Error
}
