package scanner

import (
	"util"

	"token"
)

// Scanner is responsible for lexing source files
type Scanner struct {
	input   string
	head    int
	read    int
	current rune
}

func New(input string) *Scanner {
	s := &Scanner{input: input}
	s.readChar()
	return s
}

func (s *Scanner) readChar() {
	if s.read >= len(s.input) {
		s.current = 0
	} else {
		s.current = rune(s.input[s.read])
	}

	s.head = s.read
	s.read++
}

func (s *Scanner) NextToken() token.Token {
	var tok token.Token
	s.consumeWhitespace()

	switch s.current {
	// handle brackety things
	case '[':
		tok = newToken(token.LBRACK, s.current)
	case ']':
		tok = newToken(token.RBRACK, s.current)
	case '{':
		tok = newToken(token.LBRACE, s.current)
	case '}':
		tok = newToken(token.RBRACE, s.current)
	case '(':
		tok = newToken(token.LPAREN, s.current)
	case ')':
		tok = newToken(token.RPAREN, s.current)
	case '<':
		tok = newToken(token.LCHEV, s.current)
	case '>':
		tok = newToken(token.RCHEV, s.current)
	case 0:
		tok.Literal = ""
		tok.TokenKind = token.EOF

	default:
		if util.IsLetter(s.current) {
			tok.Literal = s.scanText()
			tok.TokenKind = token.Lookup(tok.Literal)
		} else if util.IsDigit(s.current) {
			tok.Literal = s.scanNumber()
			tok.TokenKind = token.INT
		} else {
			tok = newToken(token.ILLEGAL, s.current)
		}
	}
	s.readChar()
	return tok
}

func (s *Scanner) consumeWhitespace() {
	for util.IsWhitespace(s.current) {
		s.readChar()
	}
}

// read number like things
func (s *Scanner) scanNumber() string {
	position := s.head
	for util.IsDigit(s.current) {
		s.readChar()
	}
	s.read--
	return s.input[position:s.head]
}

// read until next whitespace
func (s *Scanner) scanText() string {
	position := s.head
	for util.IsIdentifierChar(s.current) {
		s.readChar()
	}
	s.read--
	return s.input[position:s.head]
}

func (s *Scanner) peekChar() byte {
	if s.read >= len(s.input) {
		return 0
	} else {
		return s.input[s.read]
	}
}

func (s *Scanner) Source() string {
	return s.input
}

func newToken(kind token.TokenKind, ch rune) token.Token {
	return token.Token{TokenKind: kind, Literal: string(ch)}
}
