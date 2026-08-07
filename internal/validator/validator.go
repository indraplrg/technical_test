package validator

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidateErrors converts a binding/validation error into human readable
// messages using each field's name in snake_case, e.g.
// "nama_jurusan is required".
func ValidateErrors(err error) []string {
	if err == nil {
		return nil
	}

	var validationErrs validator.ValidationErrors
	if !errors.As(err, &validationErrs) {
		return []string{err.Error()}
	}

	messages := make([]string, 0, len(validationErrs))
	for _, fe := range validationErrs {
		fieldName := snakeCase(fe.Field())
		messages = append(messages, messageFor(fieldName, fe.Tag(), fe.Param()))
	}
	return messages
}

func messageFor(field, tag, param string) string {
	switch tag {
	case "required":
		return field + " is required"
	case "gt":
		return field + " must be greater than " + param
	case "gte":
		return field + " must be greater than or equal " + param
	case "email":
		return field + " must be a valid email"
	case "uuid":
		return field + " must be a valid uuid"
	case "oneof":
		return field + " must be one of: " + param
	case "date":
		return field + " must use YYYY-MM-DD format"
	default:
		return field + " is invalid"
	}
}

func snakeCase(s string) string {
	var b strings.Builder
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				b.WriteByte('_')
			}
			b.WriteRune(r + 32)
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}
