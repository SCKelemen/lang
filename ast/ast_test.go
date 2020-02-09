package ast

import (
	"testing"

	"token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&TypeDeclarationStatement{
				Token: token.Token{TokenKind: token.TYPE, Literal: "type"},
				Name: &Identifier{
					Token: token.Token{TokenKind: token.IDENT, Literal: "rwc"},
					Value: "rwc",
				},
				Value: &Identifier{
					Token: token.Token{TokenKind: token.IDENT, Literal: "ReaderWriterCloser"},
					Value: "ReaderWriterCloser",
				},
			},
		},
	}
	if program.String() != "type rwc = ReaderWriterCloser;" {
		t.Errorf("program.String() was not correct, received %q", program.String())
	}
}
