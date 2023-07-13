package gval

import "github.com/XQ-Gang/gg/constraints"

// Zero returns the zero value of the given type.
func Zero[T any]() (v T) {
	return
}

func Abs[T constraints.Real](v T) T {
	if v < 0 {
		return -v
	}
	return v
}
