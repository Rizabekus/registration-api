package validators

import (
	"time"

	"github.com/go-playground/validator"
)

func DateOfBirthValidator(fl validator.FieldLevel) bool {
	dob := fl.Field().Interface().(time.Time)

	current := time.Now()

	return dob.Before(current)
}
