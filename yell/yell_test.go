package yell

import "testing"

func YellTest(t *testing.T) {
	message := Yell("hello")
	expectation := "HELLO"

	if message != expectation {
		t.Fatalf("expected: %s\treceived: %s\n", expectation, message)
	}

}
