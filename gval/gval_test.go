package gval

import (
	"testing"

	. "github.com/XQ-Gang/gg/assert"
)

func TestZero(t *testing.T) {
	Eq(t, false, Zero[bool]())
	Eq(t, 0, Zero[int]())
	Eq(t, 0, Zero[float64]())
	Eq(t, "", Zero[string]())
	Eq(t, nil, Zero[any]())
	Eq(t, nil, Zero[*int]())
	Eq(t, nil, Zero[[]int]())
	Eq(t, nil, Zero[map[int]int]())
	type MyStruct struct{ A int }
	Eq(t, MyStruct{}, Zero[MyStruct]())
	Eq(t, nil, Zero[*MyStruct]())
}

func TestAbs(t *testing.T) {
	Eq(t, 0, Abs(0))
	Eq(t, 1, Abs(1))
	Eq(t, 1, Abs(-1))
	Eq(t, 1.5, Abs(-1.5))
	Eq(t, 1.5, Abs(1.5))
	Eq(t, 8, Abs(uint(8)))
}
