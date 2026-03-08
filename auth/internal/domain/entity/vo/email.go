package vo

import (
	"errors"
	"regexp"
)

var (
	InvalidEmailError    = errors.New("invalid_email_error")
	MainEmailLengthError = errors.New("min_email_length_error")
	MaxEmailLengthError  = errors.New("max_email_length_error")
)

type Email struct {
	value string
}

func NewEmail(plainEmail string) (Email, error) {
	if len(plainEmail) < MinEmailLength {
		return Email{}, MainEmailLengthError

	}
	if len(plainEmail) > MaxEmailLength {
		return Email{}, MaxEmailLengthError

	}
	if !regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`).MatchString(plainEmail) {
		return Email{}, InvalidEmailError
	}
	return Email{value: plainEmail}, nil
}

func (email *Email) String() string {
	return email.value
}
