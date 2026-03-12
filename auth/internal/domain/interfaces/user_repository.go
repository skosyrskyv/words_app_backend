package interfaces

import "auth/internal/domain/entity"

type UserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserByUUID(uuid string) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	UpdateUserPassword(user *entity.User) error
	DeleteUser(uuid string) error
}
