package parser

import "testing"

func SanityTest(t *testing.T) {
	message := SanityCheck("hello")
	expectation := "checked hello"

	if message != expectation {
		t.Fatalf("expected: %s\treceived: %s\n", expectation, message)
	}

}
