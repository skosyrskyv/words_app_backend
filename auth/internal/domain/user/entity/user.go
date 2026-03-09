package entity

import (
	"auth/internal/domain/user/entity/vo"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UUID         uuid.UUID
	Email        vo.Email
	PasswordHash vo.PasswordHash
	CreatedAt    time.Time
	UpdatedAt    time.Time
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

	now := time.Now()

	return &User{
		UUID:         uuid.New(),
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}

func (u *User) VerifyPassword(plainPassword string) (bool, error) {
	password, err := vo.NewPassword(plainPassword)
	if err != nil {
		return false, err
	}

	isValid := bcrypt.CompareHashAndPassword(
		[]byte(u.PasswordHash.String()),
		[]byte(password.String()),
	)

	return isValid == nil, nil
}

func (u *User) ChangePassword(newPassword string) error {
	plainPassword, err := vo.NewPassword(newPassword)
	if err != nil {
		return err
	}

	passwordHash, err := vo.NewPasswordHash(plainPassword)
	if err != nil {
		return err
	}

	u.PasswordHash = passwordHash

	return nil
}

func (u *User) GetEmail() string {
	return u.Email.String()
}

func (u *User) GetPasswordHash() string {
	return u.PasswordHash.String()
}
