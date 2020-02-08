package parser

import (
	"scanner"
	"token"
)

// SanityCheck checks your sanity
func SanityCheck(message string) string {
	return "checked " + message
}

type Parser struct {
	lxr          *scanner.Scanner
	currentToken token.Token
	peekToken    token.Token

	errors []string
}
