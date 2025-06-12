package utils

import "regexp"

func IsSnakeCase(str string) bool {
	// Regular expression for snake_case validation
	var regex = regexp.MustCompile(`^[a-z0-9]+(_[a-z0-9]+)*$`)
	return regex.MatchString(str)
}
