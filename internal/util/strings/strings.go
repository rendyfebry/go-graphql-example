package strings

import "regexp"

// SanitizeToAlphaNumeric will sanitize a string and return alpha numeric and space only
func SanitizeToAlphaNumeric(src string) string {
	reg, _ := regexp.Compile(`[^a-zA-Z0-9\s]+`)

	return reg.ReplaceAllString(src, "")
}
