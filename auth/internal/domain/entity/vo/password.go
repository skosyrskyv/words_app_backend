package vo

import "errors"

var (
	PasswordTooShortError  = errors.New("password_too_short_error")
	PasswordTooLengthError = errors.New("password_too_short_error")
)

type Password struct {
	value string
}

func NewPassword(plainPassword string) (Password, error) {
	if len(plainPassword) < MinPasswordLength {
		return Password{}, PasswordTooShortError
	}

	if len(plainPassword) > MaxPasswordLength {
		return Password{}, PasswordTooLengthError
	}

	return Password{value: plainPassword}, nil
}

func (password *Password) String() string {
	return password.value
}
