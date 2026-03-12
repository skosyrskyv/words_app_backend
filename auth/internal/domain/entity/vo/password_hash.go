package vo

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	EmptyPasswordError = errors.New("empty_password_error")
)

type PasswordHash struct {
	value string
}

func NewPasswordHash(password Password) (PasswordHash, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password.value), bcrypt.DefaultCost)

	if err != nil {
		return PasswordHash{}, err
	}

	return PasswordHash{value: string(hashBytes)}, nil
}

func NewPasswordHashFromString(hash string) (PasswordHash, error) {
	if hash == "" {
		return PasswordHash{}, EmptyPasswordError
	}
	return PasswordHash{value: hash}, nil
}

func (hash *PasswordHash) String() string {
	return hash.value
}
