package util

import (
	"testing"
)

func TestDotNotIdentifierChar(t *testing.T) {
	if IsIdentifierChar('.') {
		t.Fatalf("Character '.' is an identifier char, and should not be")
	}
}
