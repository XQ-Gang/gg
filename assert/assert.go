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
