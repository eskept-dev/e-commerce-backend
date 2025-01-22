package errors

import "errors"

var (
	// Authentication errors
	ErrInvalidCredentials = errors.New("INVALID_CREDENTIALS")
	ErrEmailExists        = errors.New("EMAIL_ALREADY_EXISTS")
	ErrUserNotEnabled     = errors.New("USER_NOT_ENABLED")

	// Authorization errors
	ErrInvalidToken = errors.New("INVALID_TOKEN")
	ErrTokenExpired = errors.New("TOKEN_EXPIRED")

	// Database errors
	ErrNotFound = errors.New("NOT_FOUND")

	// Common errors
	ErrInternalServerError = errors.New("INTERNAL_SERVER_ERROR")
	ErrAlreadyExists       = errors.New("ALREADY_EXISTS")

	// Validation errors
	ErrInvalidRequest = errors.New("INVALID_REQUEST")

	// User Profile errors
	ErrInvalidProfile = errors.New("INVALID_PROFILE")
)
