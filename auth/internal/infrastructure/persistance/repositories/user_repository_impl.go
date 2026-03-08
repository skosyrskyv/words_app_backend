package user_repository_impl

import (
	"auth/internal/domain/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (repository *userRepository) GetByUUID(uuid string) (*entity.User, error) {
	var user entity.User
	err := repository.db.First(&user, uuid).Error
	return &user, err
}

func (r *userRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

// func (repository *userRepository) UpdatePassword(user *entity.User) error {
// 	return repository.db.Model(&user).Update("password", user.Password).Error
// }

// func (repository *userRepository) Delete(uuid string) error {
// 	return repository.db.Delete(entity.User).Error
// }
