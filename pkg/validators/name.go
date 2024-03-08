package validators

import (
	"unicode"

	"github.com/go-playground/validator"
)

func CyrillicOrLatinValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	hasCyrillic := false
	hasLatin := false
	for _, char := range value {
		if unicode.Is(unicode.Cyrillic, char) {
			hasCyrillic = true
		}
		if unicode.Is(unicode.Latin, char) {
			hasLatin = true
		}
		if hasCyrillic && hasLatin {
			return false
		}
	}
	return hasCyrillic || hasLatin
}
