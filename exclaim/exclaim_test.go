package exclaim

import "testing"

func ExclaimsTest(t *testing.T) {
	message := Exclaim("hello")
	expectation := "HELLO!!!"

	if message != expectation {
		t.Fatalf("expected: %s\treceived: %s\n", expectation, message)
	}

}
