package utils

import (
	"regexp"
	"strings"
)

//var keywords = map[string]int8{"ID": 1, "JWT": 1, "URL": 1}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToUpperSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToUpper(snake)
}

func EnvToViperKey(str string) string {
	dot := strings.ReplaceAll(str, "_", ".")
	return strings.ToLower(dot)
}

// DerefString takes a pointer to string and the corresponding value or empty string if pointer is nil
func DerefString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
