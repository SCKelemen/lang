package token

import (
	"strconv"
)

type TokenKind int

type Token struct {
	TokenKind TokenKind
	Literal   string
}

const (
	ILLEGAL TokenKind = iota
	EOF

	IDENT
	INT // for natural numbers

	LBRACK // [
	RBRACK // ]
	LBRACE // {
	RBRACE // }
	LPAREN // (
	RPAREN // )
	LCHEV  // <
	RCHEV  // >

	COMMA // ,
	DOT   // .
	COLON // :
	SEMI  // ;

	ASSIGN // ==

	PIPE // |
	AMP  // &
	BANG // !

	FPIPE // |>
	RPIPE // <|

	RANGE  // ..
	SPREAD // ...

	// arithmeticy bits
	NEG // -
	SUM // +
	MUL // *
	QUO // /

	EQL  // ==
	NEQL // !=

	_keywords_beg
	TYPE
	MATCH
	RETURN
	TRUE
	FALSE
	FUNC
	LET
	_keywords_end
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	IDENT: "IDENTITY",
	INT:   "INT",

	LBRACK: "[",
	RBRACK: "]",
	LBRACE: "{",
	RBRACE: "}",
	LPAREN: "(",
	RPAREN: ")",
	LCHEV:  "<",
	RCHEV:  ">",

	COMMA: ",",
	DOT:   ".",
	COLON: ":",
	SEMI:  ";",

	ASSIGN: "=",

	PIPE: "|",
	AMP:  "&",
	BANG: "!",

	FPIPE: "|>",
	RPIPE: "<|",

	RANGE:  "..",
	SPREAD: "...",

	NEG: "-",
	SUM: "+",
	MUL: "*",
	QUO: "/",

	EQL:  "==",
	NEQL: "!=",

	TYPE:   "type",
	MATCH:  "match",
	RETURN: "return",
	TRUE:   "true",
	FALSE:  "false",
	FUNC:   "func",
	LET:    "let",
}

func (token TokenKind) String() string {
	s := ""
	if 0 <= token && token < TokenKind(len(tokens)) {
		s = tokens[token]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(token)) + ")"
	}
	return s
}

var keywords map[string]TokenKind

func init() {
	keywords = make(map[string]TokenKind)
	for i := _keywords_beg + 1; i < _keywords_end; i++ {
		keywords[tokens[i]] = i
	}
}

func Lookup(candidate string) TokenKind {
	if tok, isKeyword := keywords[candidate]; isKeyword {
		return tok
	}
	return IDENT
}

func SanityCheck() string {
	return "token"
}
