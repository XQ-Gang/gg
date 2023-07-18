package assert

import (
	"testing"
)

func TestEq(t *testing.T) {
	Eq(t, 1, 1)
	Eq(t, 1.1, 1.1)
	Eq(t, "1", "1")
	Eq(t, any(nil), any(nil))
	a, b := 1, 1
	Eq(t, []*int{&a}, []*int{&b})
	Eq(t, map[int]int{1: 1}, map[int]int{1: 1})
}

func TestPanic(t *testing.T) {
	Panic(t, func() { panic("panic") })
}
