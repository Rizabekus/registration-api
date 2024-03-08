package validators

import (
	"regexp"

	"github.com/go-playground/validator"
)

func PhoneValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	phoneRegex := regexp.MustCompile(`^\+[1-9]\d{1,14}$`)

	return phoneRegex.MatchString(value)
}
