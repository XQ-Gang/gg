package gutil

import (
	"testing"

	. "github.com/XQ-Gang/gg/assert"
)

func TestIf(t *testing.T) {
	Eq(t, If(true, 1), 1)
	Eq(t, If(true, 1, 2), 1)
	Eq(t, If(false, 1), 0)
	Eq(t, If(false, 1, 2), 2)
}

func TestIfF(t *testing.T) {
	ifFunc, elseFunc := func() int { return 1 }, func() int { return 2 }
	Eq(t, IfF(true, ifFunc), 1)
	Eq(t, IfF(true, ifFunc, elseFunc), 1)
	Eq(t, IfF(false, ifFunc), 0)
	Eq(t, IfF(false, ifFunc, elseFunc), 2)
}

func TestIfFL(t *testing.T) {
	ifFunc := func() int { return 1 }
	Eq(t, IfFL(true, ifFunc, 2), 1)
	Eq(t, IfFL(false, ifFunc, 2), 2)
}

func TestIfFR(t *testing.T) {
	elseFunc := func() int { return 2 }
	Eq(t, IfFR(true, 1, elseFunc), 1)
	Eq(t, IfFR(false, 1, elseFunc), 2)
}
