package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"unicode"
)

// ToCamelCase converts a string to camelCase.
func ToCamelCase(str string) string {
	words := strings.FieldsFunc(str, func(r rune) bool {
		// Split string into words on spaces and punctuation
		return unicode.IsSpace(r) || unicode.IsPunct(r)
	})

	for i := 0; i < len(words); i++ {
		// catch PascalCase and return the camelCase equivalent
		if len(words) == 1 {
			r := []rune(words[i])
			r[0] = unicode.ToLower(r[0])
			return string(r)
		}
		if i == 0 {
			// Make the first word all lower-case
			words[i] = strings.ToLower(words[i])
			continue
		}
		// Make the rest of the words title case
		words[i] = cases.Title(language.English, cases.Compact).String(words[i])
	}

	// Join words together
	return strings.Join(words, "")
}
