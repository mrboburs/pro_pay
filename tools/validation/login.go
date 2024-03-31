package validation

import (
	"regexp"
)

// IsValidLogin ...
func IsValidLogin(login string) bool {
	r := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]$`)
	return r.MatchString(login)
}
