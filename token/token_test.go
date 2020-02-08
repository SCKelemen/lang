package token

import "testing"

func TokeSanityTest(t *testing.T) {
	message := SanityCheck()
	expectation := "token"

	if message != expectation {
		t.Fatalf("expected: %s\treceived: %s\n", expectation, message)
	}

}
