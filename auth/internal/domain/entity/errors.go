package entity

import "errors"

var (
	ErrUserNotFound       = errors.New("user_not_found_error")
	ErrEmailAlreadyExists = errors.New("email_already_exists_error")
	ErrUserCredentials    = errors.New("user_credentials_error")
)
