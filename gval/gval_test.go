package gval

import (
	"testing"
)

func TestAbs(t *testing.T) {
	t.Log(Abs(-1))
	t.Log(Abs(0))
	t.Log(Abs(1))

	t.Log(Abs(-1.234))
	t.Log(Abs(0.0))
	t.Log(Abs(1.234))

}
