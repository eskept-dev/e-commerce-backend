package errors

import "errors"

var (
	// User-related errors
	ErrEmailExists        = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrAccountNotActive   = errors.New("account is not activated")
	ErrUserNotFound       = errors.New("user not found")
	ErrOldPasswordInvalid = errors.New("old password is incorrect")

	// General errors
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrInternal     = errors.New("internal server error")
)
