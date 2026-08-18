package domain

import "fmt"

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string { return fmt.Sprintf("%s: %s", e.Field, e.Message) }
func RequireString(field, value string) error {
	if value == "" {
		return ValidationError{field, "is required"}
	}
	return nil
}
func RequirePositive(field string, value int64) error {
	if value <= 0 {
		return ValidationError{field, "must be positive"}
	}
	return nil
}
func RequireVersion(value int64) error { return RequirePositive("version", value) }
func ValidateLines(lines []OrderLine) error {
	if len(lines) == 0 {
		return ValidationError{"lines", "must not be empty"}
	}
	for _, line := range lines {
		if err := RequireString("sku", line.SKU); err != nil {
			return err
		}
		if err := RequirePositive("quantity", line.Quantity); err != nil {
			return err
		}
	}
	return nil
}
