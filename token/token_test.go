package token

import "testing"

func ScreamssTest(t *testing.T) {
	message := Token("hello")
	expectation := "HELLO!!! +Token"

	if message != expectation {
		t.Fatalf("expected: %s\treceived: %s\n", expectation, message)
	}

}
