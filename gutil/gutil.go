package gutil

import "github.com/XQ-Gang/gg/gval"

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

func IfF[T any](cond bool, ifFunc IfFunc[T], elseFunc ...IfFunc[T]) T {
	if cond {
		return ifFunc()
	}
	if len(elseFunc) > 0 {
		return elseFunc[0]()
	}
	return gval.Zero[T]()
}

func IfFL[T any](cond bool, ifFunc IfFunc[T], elseVal T) T {
	if cond {
		return ifFunc()
	}
	return elseVal
}

func IfFR[T any](cond bool, ifVal T, elseFunc IfFunc[T]) T {
	if cond {
		return ifVal
	}
	return elseFunc()
}
