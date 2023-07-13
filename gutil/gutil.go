package gutil

import "github.com/XQ-Gang/gg/gval"

// If returns ifVal if cond is true, elseVal otherwise.
// If cond is true, ifVal is returned, elseVal is ignored.
// If cond is false and elseVal is provided, elseVal is returned.
// If cond is false and elseVal is not provided, the zero value of T is returned.
func If[T any](cond bool, ifVal T, elseVal ...T) T {
	if cond {
		return ifVal
	}
	if len(elseVal) > 0 {
		return elseVal[0]
	}
	return gval.Zero[T]()
}

type IfFunc[T any] func() T

// IfF returns the result of ifFunc if cond is true, the result of elseFunc otherwise.
// If cond is true, ifFunc is called, elseFunc is ignored.
// If cond is false and elseFunc is provided, elseFunc is called.
// If cond is false and elseFunc is not provided, the zero value of T is returned.
func IfF[T any](cond bool, ifFunc IfFunc[T], elseFunc ...IfFunc[T]) T {
	if cond {
		return ifFunc()
	}
	if len(elseFunc) > 0 {
		return elseFunc[0]()
	}
	return gval.Zero[T]()
}

// IfFL returns the result of ifFunc if cond is true, elseVal otherwise.
// If cond is true, ifFunc is called, elseVal is ignored.
// If cond is false, elseVal is returned.
func IfFL[T any](cond bool, ifFunc IfFunc[T], elseVal T) T {
	if cond {
		return ifFunc()
	}
	return elseVal
}

// IfFR returns ifVal if cond is true, the result of elseFunc otherwise.
// If cond is true, ifVal is returned, elseFunc is ignored.
// If cond is false, elseFunc is called.
func IfFR[T any](cond bool, ifVal T, elseFunc IfFunc[T]) T {
	if cond {
		return ifVal
	}
	return elseFunc()
}
