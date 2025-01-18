package errors

import "errors"

var (
	// Authentication errors
	ErrInvalidCredentials = errors.New("INVALID_CREDENTIALS")
	ErrEmailExists        = errors.New("EMAIL_ALREADY_EXISTS")

	// Database errors
	ErrUserNotFound = errors.New("USER_NOT_FOUND")

	// Common errors
	ErrInternalServerError = errors.New("INTERNAL_SERVER_ERROR")

	// Validation errors
	ErrInvalidRequest = errors.New("INVALID_REQUEST")
)
