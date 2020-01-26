package yell

import "strings"

// Yell announces your message louder
func Yell(message string) string {
	return strings.ToUpper(message)
}
