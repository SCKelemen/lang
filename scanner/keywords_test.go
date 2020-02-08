package scanner

import (
	"testing"
	"token"
)

func TestKeywords(t *testing.T) {
	input := `
type
match
return
true
false
func
let 
	`
	tests := []struct {
		expectedKind    token.TokenKind
		expectedLiteral string
	}{
		{token.TYPE, "type"},
		{token.MATCH, "match"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.FALSE, "false"},
		{token.FUNC, "func"},
		{token.LET, "let"},
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

func TestIdentities(t *testing.T) {
	input := `
id_type
id_match
id_return
id_true
id_false
id_func
id_let 
	`
	tests := []struct {
		expectedKind    token.TokenKind
		expectedLiteral string
	}{
		{token.IDENT, "id_type"},
		{token.IDENT, "id_match"},
		{token.IDENT, "id_return"},
		{token.IDENT, "id_true"},
		{token.IDENT, "id_false"},
		{token.IDENT, "id_func"},
		{token.IDENT, "id_let"},
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
