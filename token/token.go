package token

import "strings"

// Token yells your message louder
func Token(message string) string {
	return strings.ToUpper(message) + "!!! +Token"
}
