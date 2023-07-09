package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"regexp"
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

// ToPascalCase converts a string to PascalCase.
func ToPascalCase(str string) string {
	words := strings.FieldsFunc(str, func(r rune) bool {
		// Split string into words on spaces and punctuation
		return unicode.IsSpace(r) || unicode.IsPunct(r)
	})

	for i := 0; i < len(words); i++ {
		// Make each word title case
		words[i] = cases.Title(language.English, cases.Compact).String(words[i])
	}

	// Join words together
	return strings.Join(words, "")
}

func PathToTitle(s string) string {
	parts := strings.Split(s, "/")
	result := ""

	for i := range parts {
		if parts[i] != "" {
			if strings.HasPrefix(parts[i], "{") && strings.HasSuffix(parts[i], "}") {
				// Extract slug name, capitalise it, remove the braces and prefix with 'By'
				slug := parts[i][1 : len(parts[i])-1] // remove braces
				slugRune := []rune(slug)
				slugRune[0] = unicode.ToUpper(slugRune[0]) // make first letter capital
				slug = string(slugRune)
				result += "By" + slug // prefix with 'By'
				continue
			}
			r := []rune(parts[i])
			r[0] = unicode.ToUpper(r[0])
			result += string(r)
		}
	}

	return result
}

func ExtractSlugs(path string) []string {
	r := regexp.MustCompile(`\{([^}]*)\}`)
	matches := r.FindAllStringSubmatch(path, -1)

	var slugs []string
	for _, match := range matches {
		slugs = append(slugs, match[1])
	}

	return slugs
}
