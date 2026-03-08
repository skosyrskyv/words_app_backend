package repository

import "auth/internal/domain/entity"

type UserRepository interface {
	Create(user *entity.User) (string, error)
	GetByUUID(uuid string) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	UpdatePassword(user *entity.User) error
	Delete(uuid string) error
}
