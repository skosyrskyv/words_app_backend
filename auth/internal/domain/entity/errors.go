package entity

import "errors"

var (
	UserNotFoundError       = errors.New("user_not_found_error")
	EmailAlreadyExistsError = errors.New("email_already_exists_error")
)
