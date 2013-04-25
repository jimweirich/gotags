package assert

import "testing"

func StringEqual(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("\nexpected: %v\ngot     : %v", expected, actual)
	}
}
