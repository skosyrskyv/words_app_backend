package entity

import (
	"auth/internal/domain/entity/vo"

	"github.com/google/uuid"
)

type User struct {
	UUID         uuid.UUID
	Email        vo.Email
	PasswordHash vo.PasswordHash
}

func NewUser(plainEmail string, plainPassword string) (*User, error) {
	email, err := vo.NewEmail(plainEmail)

	if err != nil {
		return nil, err
	}

	password, err := vo.NewPassword(plainPassword)

	if err != nil {
		return nil, err
	}

	passwordHash, err := vo.NewPasswordHash(password)

	if err != nil {
		return nil, err
	}

	return &User{
		UUID:         uuid.New(),
		Email:        email,
		PasswordHash: passwordHash,
	}, nil
}
