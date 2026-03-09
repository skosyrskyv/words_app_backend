package repository

import "auth/internal/domain/user/entity"

type Postgres interface {
	Create(user *entity.User) (*entity.User, error)
	GetByUUID(uuid string) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	UpdatePassword(user *entity.User) error
	Delete(uuid string) error
}
