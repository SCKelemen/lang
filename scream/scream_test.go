package scream

import "testing"

func ScreamssTest(t *testing.T) {
	message := Scream("hello")
	expectation := "HELLO!!!"

	if message != expectation {
		t.Fatalf("expected: %s\treceived: %s\n", expectation, message)
	}

}
