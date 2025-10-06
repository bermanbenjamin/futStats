package validation

import (
	"testing"
)

// Test cases for email validation (pure function testing)
func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{
			name:     "valid email",
			email:    "test@example.com",
			expected: true,
		},
		{
			name:     "valid email with subdomain",
			email:    "user@mail.example.com",
			expected: true,
		},
		{
			name:     "empty email",
			email:    "",
			expected: false,
		},
		{
			name:     "invalid format",
			email:    "not-an-email",
			expected: false,
		},
		{
			name:     "missing domain",
			email:    "test@",
			expected: false,
		},
		{
			name:     "missing local part",
			email:    "@example.com",
			expected: false,
		},
		{
			name:     "too long email",
			email:    string(make([]byte, 255)) + "@example.com",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateEmail(tt.email)
			if result.IsValid != tt.expected {
				t.Errorf("ValidateEmail(%q) = %v, expected %v", tt.email, result.IsValid, tt.expected)
			}
		})
	}
}

// Test cases for password validation (pure function testing)
func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		expected bool
	}{
		{
			name:     "valid password",
			password: "Password123",
			expected: true,
		},
		{
			name:     "empty password",
			password: "",
			expected: false,
		},
		{
			name:     "too short",
			password: "Pass1",
			expected: false,
		},
		{
			name:     "no uppercase",
			password: "password123",
			expected: false,
		},
		{
			name:     "no lowercase",
			password: "PASSWORD123",
			expected: false,
		},
		{
			name:     "no numbers",
			password: "Password",
			expected: false,
		},
		{
			name:     "too long",
			password: string(make([]byte, 129)),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidatePassword(tt.password)
			if result.IsValid != tt.expected {
				t.Errorf("ValidatePassword(%q) = %v, expected %v", tt.password, result.IsValid, tt.expected)
			}
		})
	}
}

// Test cases for name validation (pure function testing)
func TestValidateName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "valid name",
			input:    "John Doe",
			expected: true,
		},
		{
			name:     "valid name with hyphen",
			input:    "Jean-Pierre",
			expected: true,
		},
		{
			name:     "valid name with apostrophe",
			input:    "O'Connor",
			expected: true,
		},
		{
			name:     "empty name",
			input:    "",
			expected: false,
		},
		{
			name:     "too short",
			input:    "A",
			expected: false,
		},
		{
			name:     "too long",
			input:    string(make([]byte, 101)),
			expected: false,
		},
		{
			name:     "invalid characters",
			input:    "John123",
			expected: false,
		},
		{
			name:     "whitespace only",
			input:    "   ",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateName(tt.input)
			if result.IsValid != tt.expected {
				t.Errorf("ValidateName(%q) = %v, expected %v", tt.input, result.IsValid, tt.expected)
			}
		})
	}
}

// Test cases for age validation (pure function testing)
func TestValidateAge(t *testing.T) {
	tests := []struct {
		name     string
		age      int
		expected bool
	}{
		{
			name:     "valid age",
			age:      25,
			expected: true,
		},
		{
			name:     "minimum age",
			age:      16,
			expected: true,
		},
		{
			name:     "maximum age",
			age:      100,
			expected: true,
		},
		{
			name:     "too young",
			age:      15,
			expected: false,
		},
		{
			name:     "too old",
			age:      101,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateAge(tt.age)
			if result.IsValid != tt.expected {
				t.Errorf("ValidateAge(%d) = %v, expected %v", tt.age, result.IsValid, tt.expected)
			}
		})
	}
}

// Test cases for complete player validation (pure function testing)
func TestValidatePlayer(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		password string
		age      int
		expected bool
	}{
		{
			name:     "valid player",
			email:    "john@example.com",
			password: "Password123",
			age:      25,
			expected: true,
		},
		{
			name:     "invalid name",
			email:    "john@example.com",
			password: "Password123",
			age:      25,
			expected: false,
		},
		{
			name:     "invalid email",
			email:    "invalid-email",
			password: "Password123",
			age:      25,
			expected: false,
		},
		{
			name:     "invalid password",
			email:    "john@example.com",
			password: "weak",
			age:      25,
			expected: false,
		},
		{
			name:     "invalid age",
			email:    "john@example.com",
			password: "Password123",
			age:      15,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidatePlayer(tt.name, tt.email, tt.password, tt.age)
			if result.IsValid != tt.expected {
				t.Errorf("ValidatePlayer() = %v, expected %v", result.IsValid, tt.expected)
			}
		})
	}
}
