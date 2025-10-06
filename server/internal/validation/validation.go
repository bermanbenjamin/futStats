package validation

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

// ValidationError represents a validation error with field and message
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidationResult contains validation results
type ValidationResult struct {
	IsValid bool              `json:"is_valid"`
	Errors  []ValidationError `json:"errors,omitempty"`
}

// ValidateEmail validates email format (pure function)
func ValidateEmail(email string) ValidationResult {
	if email == "" {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{{
				Field:   "email",
				Message: "email is required",
			}},
		}
	}

	if len(email) > 254 {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{{
				Field:   "email",
				Message: "email is too long (max 254 characters)",
			}},
		}
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{{
				Field:   "email",
				Message: "invalid email format",
			}},
		}
	}

	return ValidationResult{IsValid: true}
}

// ValidatePassword validates password strength (pure function)
func ValidatePassword(password string) ValidationResult {
	var errors []ValidationError

	if password == "" {
		errors = append(errors, ValidationError{
			Field:   "password",
			Message: "password is required",
		})
		return ValidationResult{IsValid: false, Errors: errors}
	}

	if len(password) < 8 {
		errors = append(errors, ValidationError{
			Field:   "password",
			Message: "password must be at least 8 characters long",
		})
	}

	if len(password) > 128 {
		errors = append(errors, ValidationError{
			Field:   "password",
			Message: "password is too long (max 128 characters)",
		})
	}

	// Check for at least one uppercase letter
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	if !hasUpper {
		errors = append(errors, ValidationError{
			Field:   "password",
			Message: "password must contain at least one uppercase letter",
		})
	}

	// Check for at least one lowercase letter
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	if !hasLower {
		errors = append(errors, ValidationError{
			Field:   "password",
			Message: "password must contain at least one lowercase letter",
		})
	}

	// Check for at least one number
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	if !hasNumber {
		errors = append(errors, ValidationError{
			Field:   "password",
			Message: "password must contain at least one number",
		})
	}

	return ValidationResult{
		IsValid: len(errors) == 0,
		Errors:  errors,
	}
}

// ValidateName validates player name (pure function)
func ValidateName(name string) ValidationResult {
	if name == "" {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{{
				Field:   "name",
				Message: "name is required",
			}},
		}
	}

	name = strings.TrimSpace(name)
	if len(name) < 2 {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{{
				Field:   "name",
				Message: "name must be at least 2 characters long",
			}},
		}
	}

	if len(name) > 100 {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{{
				Field:   "name",
				Message: "name is too long (max 100 characters)",
			}},
		}
	}

	// Check for valid characters (letters, spaces, hyphens, apostrophes)
	validName := regexp.MustCompile(`^[a-zA-Z\s\-']+$`).MatchString(name)
	if !validName {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{{
				Field:   "name",
				Message: "name can only contain letters, spaces, hyphens, and apostrophes",
			}},
		}
	}

	return ValidationResult{IsValid: true}
}

// ValidateUUID validates UUID format (pure function)
func ValidateUUID(id string) ValidationResult {
	if id == "" {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{{
				Field:   "id",
				Message: "id is required",
			}},
		}
	}

	if _, err := uuid.Parse(id); err != nil {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{{
				Field:   "id",
				Message: "invalid UUID format",
			}},
		}
	}

	return ValidationResult{IsValid: true}
}

// ValidateAge validates age range (pure function)
func ValidateAge(age int) ValidationResult {
	if age < 16 {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{{
				Field:   "age",
				Message: "age must be at least 16",
			}},
		}
	}

	if age > 100 {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{{
				Field:   "age",
				Message: "age must be at most 100",
			}},
		}
	}

	return ValidationResult{IsValid: true}
}

// ValidatePlayer validates all player fields (pure function)
func ValidatePlayer(name, email, password string, age int) ValidationResult {
	var allErrors []ValidationError

	// Validate name
	if nameResult := ValidateName(name); !nameResult.IsValid {
		allErrors = append(allErrors, nameResult.Errors...)
	}

	// Validate email
	if emailResult := ValidateEmail(email); !emailResult.IsValid {
		allErrors = append(allErrors, emailResult.Errors...)
	}

	// Validate password
	if passwordResult := ValidatePassword(password); !passwordResult.IsValid {
		allErrors = append(allErrors, passwordResult.Errors...)
	}

	// Validate age
	if ageResult := ValidateAge(age); !ageResult.IsValid {
		allErrors = append(allErrors, ageResult.Errors...)
	}

	return ValidationResult{
		IsValid: len(allErrors) == 0,
		Errors:  allErrors,
	}
}
