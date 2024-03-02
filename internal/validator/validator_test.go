package validator

import (
	"slices"
	"testing"

	"github.com/vbrenister/lets-go-snippetbox/internal/assert"
)

func TestAddFieldError(t *testing.T) {
	field := "test"
	fieldError := "failed"

	var v Validator

	v.AddFieldError(field, fieldError)

	actual, ok := v.FieldErrors[field]

	assert.True(t, ok)
	assert.Equal(t, actual, fieldError)
}

func TestAddNonFieldError(t *testing.T) {
	errorMessage := "failed"

	var v Validator

	v.AddNonFieldError(errorMessage)

	if !slices.Contains(v.NonFieldErrors, errorMessage) {
		t.Errorf("'%v' is not contained in %v", errorMessage, v.NonFieldErrors)
	}
}

func TestValid(t *testing.T) {
	var v Validator

	assert.True(t, v.Valid())

	v.AddFieldError("test", "failed")

	assert.False(t, v.Valid())
}

func TestCheckField(t *testing.T) {
	field := "test"
	fieldError := "failed"

	var v Validator

	v.CheckField(false, field, fieldError)

	assert.False(t, v.Valid())

	actual, ok := v.FieldErrors[field]

	assert.True(t, ok)
	assert.Equal(t, actual, fieldError)
}

func TestNotBlank(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected bool
	}{
		{
			name:     "Empty",
			value:    "",
			expected: false,
		},
		{
			name:     "Not Empty",
			value:    "test",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, NotBlank(tt.value), tt.expected)
		})
	}
}

func TestMaxChars(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		maxChars int
		expected bool
	}{
		{
			name:     "Withing range",
			value:    "test",
			maxChars: 10,
			expected: true,
		},
		{
			name:     "Out of range",
			value:    "test",
			maxChars: 2,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, MaxChars(tt.value, tt.maxChars), tt.expected)
		})
	}
}

func TestPermittedValue(t *testing.T) {
	tests := []struct {
		name            string
		permittedValues []string
		value           string
		expected        bool
	}{
		{
			name:            "In permitted list",
			permittedValues: []string{"one", "two"},
			value:           "one",
			expected:        true,
		},
		{
			name:            "Not in permitted list",
			permittedValues: []string{"one", "two"},
			value:           "three",
			expected:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, PermittedValue(tt.value, tt.permittedValues...), tt.expected)
		})
	}
}

func TestMinChars(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		maxChars int
		expected bool
	}{
		{
			name:     "Withing range",
			value:    "test",
			maxChars: 2,
			expected: true,
		},
		{
			name:     "Out of range",
			value:    "test",
			maxChars: 10,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, MinChars(tt.value, tt.maxChars), tt.expected)
		})
	}
}

func TestMatchesEmail(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected bool
	}{
		{
			name:     "Valid email",
			value:    "test@mail.org",
			expected: true,
		},
		{
			name:     "Invalid email",
			value:    "123",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, Matches(tt.value, EmailRX), tt.expected)
		})
	}
}
