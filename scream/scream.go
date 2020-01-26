package scream

import "strings"

// Scream yells your message louder
func Scream(message string) string {
	return strings.ToUpper(message) + "!!!"
}
