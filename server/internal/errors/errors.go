package errors

import (
	"fmt"
	"net/http"
)

// ErrorCode represents a specific error type
type ErrorCode string

const (
	// Authentication errors
	ErrCodeUnauthorized ErrorCode = "UNAUTHORIZED"
	ErrCodeForbidden    ErrorCode = "FORBIDDEN"
	ErrCodeInvalidToken ErrorCode = "INVALID_TOKEN"
	ErrCodeTokenExpired ErrorCode = "TOKEN_EXPIRED"

	// Validation errors
	ErrCodeValidationFailed ErrorCode = "VALIDATION_FAILED"
	ErrCodeInvalidInput     ErrorCode = "INVALID_INPUT"
	ErrCodeMissingField     ErrorCode = "MISSING_FIELD"

	// Resource errors
	ErrCodeNotFound      ErrorCode = "NOT_FOUND"
	ErrCodeAlreadyExists ErrorCode = "ALREADY_EXISTS"
	ErrCodeConflict      ErrorCode = "CONFLICT"

	// Server errors
	ErrCodeInternalError      ErrorCode = "INTERNAL_ERROR"
	ErrCodeServiceUnavailable ErrorCode = "SERVICE_UNAVAILABLE"
	ErrCodeDatabaseError      ErrorCode = "DATABASE_ERROR"
)

// AppError represents a structured application error
type AppError struct {
	Code       ErrorCode `json:"code"`
	Message    string    `json:"message"`
	Details    string    `json:"details,omitempty"`
	StatusCode int       `json:"-"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("%s: %s (%s)", e.Code, e.Message, e.Details)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// NewAppError creates a new application error
func NewAppError(code ErrorCode, message string, statusCode int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
	}
}

// WithDetails adds details to the error
func (e *AppError) WithDetails(details string) *AppError {
	e.Details = details
	return e
}

// Predefined errors following clean code principles

// Authentication errors
var (
	ErrUnauthorized = NewAppError(
		ErrCodeUnauthorized,
		"Authentication required",
		http.StatusUnauthorized,
	)

	ErrForbidden = NewAppError(
		ErrCodeForbidden,
		"Access denied",
		http.StatusForbidden,
	)

	ErrInvalidToken = NewAppError(
		ErrCodeInvalidToken,
		"Invalid authentication token",
		http.StatusUnauthorized,
	)

	ErrTokenExpired = NewAppError(
		ErrCodeTokenExpired,
		"Authentication token has expired",
		http.StatusUnauthorized,
	)
)

// Validation errors
var (
	ErrValidationFailed = NewAppError(
		ErrCodeValidationFailed,
		"Validation failed",
		http.StatusBadRequest,
	)

	ErrInvalidInput = NewAppError(
		ErrCodeInvalidInput,
		"Invalid input provided",
		http.StatusBadRequest,
	)

	ErrMissingField = NewAppError(
		ErrCodeMissingField,
		"Required field is missing",
		http.StatusBadRequest,
	)
)

// Resource errors
var (
	ErrNotFound = NewAppError(
		ErrCodeNotFound,
		"Resource not found",
		http.StatusNotFound,
	)

	ErrAlreadyExists = NewAppError(
		ErrCodeAlreadyExists,
		"Resource already exists",
		http.StatusConflict,
	)

	ErrConflict = NewAppError(
		ErrCodeConflict,
		"Resource conflict",
		http.StatusConflict,
	)
)

// Server errors
var (
	ErrInternalError = NewAppError(
		ErrCodeInternalError,
		"Internal server error",
		http.StatusInternalServerError,
	)

	ErrServiceUnavailable = NewAppError(
		ErrCodeServiceUnavailable,
		"Service temporarily unavailable",
		http.StatusServiceUnavailable,
	)

	ErrDatabaseError = NewAppError(
		ErrCodeDatabaseError,
		"Database operation failed",
		http.StatusInternalServerError,
	)
)

// Helper functions for common error scenarios

// NewValidationError creates a validation error with details
func NewValidationError(field, message string) *AppError {
	return ErrValidationFailed.WithDetails(fmt.Sprintf("%s: %s", field, message))
}

// NewNotFoundError creates a not found error for a specific resource
func NewNotFoundError(resource, identifier string) *AppError {
	return ErrNotFound.WithDetails(fmt.Sprintf("%s with identifier '%s' not found", resource, identifier))
}

// NewConflictError creates a conflict error with details
func NewConflictError(resource, reason string) *AppError {
	return ErrConflict.WithDetails(fmt.Sprintf("%s conflict: %s", resource, reason))
}

// IsAppError checks if an error is an AppError
func IsAppError(err error) (*AppError, bool) {
	if appErr, ok := err.(*AppError); ok {
		return appErr, true
	}
	return nil, false
}
