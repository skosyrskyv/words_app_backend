package vo

import (
	"errors"
	"regexp"
)

var (
	InvalidEmailError   = errors.New("invalid_email_error")
	MinEmailLengthError = errors.New("min_email_length_error")
	MaxEmailLengthError = errors.New("max_email_length_error")
)

var emailRegex = regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`)

type Email struct {
	value string
}

func NewEmail(plainEmail string) (Email, error) {
	if len(plainEmail) < MinEmailLength {
		return Email{}, MinEmailLengthError

	}
	if len(plainEmail) > MaxEmailLength {
		return Email{}, MaxEmailLengthError

	}
	if !emailRegex.MatchString(plainEmail) {
		return Email{}, InvalidEmailError
	}
	return Email{value: plainEmail}, nil
}

func NewEmailFromString(plainEmail string) (Email, error) {
	return NewEmail(plainEmail)
}

func (email Email) String() string {
	return email.value
}
