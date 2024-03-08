package validators

import "github.com/go-playground/validator"

func ASCIIValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	for _, r := range value {
		if r > 126 || r < 32 {
			return false
		}
	}
	return true
}
