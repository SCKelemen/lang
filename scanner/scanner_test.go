package scanner

import (
	"testing"
	"token"
)

/*
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
*/

func TestNextToken(t *testing.T) {
	input := "[]{}()<>"
	tests := []struct {
		expectedKind    token.TokenKind
		expectedLiteral string
	}{
		{token.LBRACK, "["},
		{token.RBRACK, "]"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LCHEV, "<"},
		{token.RCHEV, ">"},
		{token.EOF, ""},
	}

	scnr := New(input)
	for i, tt := range tests {
		tok := scnr.NextToken()
		if tok.TokenKind != tt.expectedKind {
			t.Fatalf("tests[%d] - tokenKind wrong. expected=%q, got=%q",
				i, tt.expectedKind, tok.TokenKind)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestMoreNextToken(t *testing.T) {
	input := `
	!	
	&
	(
	)
	*
	+
	,
	-
	/
	:
	;
	>
	[

	]
	{
		
	}
	`
	tests := []struct {
		expectedKind    token.TokenKind
		expectedLiteral string
	}{
		{token.BANG, "!"},
		//{token.NEQL, "!="},
		{token.AMP, "&"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.MUL, "*"},
		{token.SUM, "+"},
		{token.COMMA, ","},
		{token.NEG, "-"},
		//{token.DOT, "."},
		//{token.RANGE, ".."},
		{token.QUO, "/"},
		{token.COLON, ":"},
		{token.SEMI, ";"},
		//{token.RPIPE, "<|"},
		//{token.LCHEV, "<"},
		//{token.EQL, "=="},
		//{token.ASSIGN, "="},
		{token.RCHEV, ">"},
		{token.LBRACK, "["},
		{token.RBRACK, "]"},
		{token.LBRACE, "{"},
		//{token.PIPE, "|"},
		//{token.FPIPE, "|>"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	scnr := New(input)
	for i, tt := range tests {
		tok := scnr.NextToken()
		if tok.TokenKind != tt.expectedKind {
			t.Fatalf("tests[%d] - tokenKind wrong. expected=%q, got=%q\t at %s",
				i, tt.expectedKind, tok.TokenKind, scnr.InspectPosition())
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q\t at %s",
				i, tt.expectedLiteral, tok.Literal, scnr.InspectPosition())
		}
	}
}
