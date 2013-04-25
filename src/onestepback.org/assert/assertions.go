package assert

import "testing"

func True(t *testing.T, condition bool) {
	if ! condition {
		t.Errorf("expected condition to be true, but was false")
	}
}

func StringEqual(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("\nexpected: %v\ngot     : %v", expected, actual)
	}
}

func IntEqual(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("\nexpected: %v\ngot     : %v", expected, actual)
	}
}
