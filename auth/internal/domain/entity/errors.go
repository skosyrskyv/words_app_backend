package entity

import "errors"

var (
	// USER ERRORS
	ErrUserNotFound       = errors.New("user_not_found_error")
	ErrEmailAlreadyExists = errors.New("email_already_exists_error")
	ErrUserCredentials    = errors.New("user_credentials_error")

	// TOKEN ERRORS
	ErrTokenNotFound = errors.New("token_not_found_error")
	ErrInvalidToken  = errors.New("invalid_token_error")
)
