package assert

import (
	"reflect"
	"testing"
)

// Eq asserts that expected and actual are equal.
func Eq[T any](t *testing.T, expected, actual T) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Not Equal.\n"+
			"Expected: %v\n"+
			"Actual: %v",
			expected, actual)
	}
}

func Panic(t *testing.T, f func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Errorf("Expected panic, but got nothing.")
		}
	}()
	f()
}
