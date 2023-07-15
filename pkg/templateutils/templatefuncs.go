package templateutils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func ParseEndpoint(endpoint string) string {
	parts := strings.Split(endpoint, "/")

	handler := ""
	for _, part := range parts {
		if part != "" {
			if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
				handler += cases.Title(language.English, cases.NoLower).String(part[1 : len(part)-1])
				continue
			}
			handler += cases.Title(language.English, cases.NoLower).String(part)
		}
	}

	return handler + "Handler"
}
