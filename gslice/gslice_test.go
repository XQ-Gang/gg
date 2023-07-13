package gslice

import (
	"fmt"
	"github.com/XQ-Gang/gg/constraints"
	"github.com/XQ-Gang/gg/gutil"
	"github.com/XQ-Gang/gg/gval"
	"testing"
)

func Range[T constraints.Integer](start, end T, step int) []T {
	if step == 0 || start == end || (start < end && step < 0) || (start > end && step > 0) {
		return []T{}
	}
	asc := start < end && step > 0
	stepT := T(step)
	resLen := gutil.If(asc, (end-start+stepT-1)/stepT, gval.Abs((end-start+stepT+1)/stepT))
	fmt.Println(resLen)
	res := make([]T, 0, resLen)
	for i := start; gutil.If(asc, i < end, i > end); i += stepT {
		res = append(res, i)
	}
	return res
}

func TestRange(t *testing.T) {
	fmt.Println(Range(0, 10, 1))
	fmt.Println(Range(0, 10, 2))
	fmt.Println(Range(0, 10, 3))
	fmt.Println(Range(0, 10, 4))
	fmt.Println(Range(0, 10, 5))
	fmt.Println(Range(0, 10, 6))
	fmt.Println(Range(0, 10, 7))
	fmt.Println(Range(0, 10, 8))
	fmt.Println(Range(0, 10, 9))
	fmt.Println(Range(0, 10, 10))
	fmt.Println(Range(0, 10, 11))
	fmt.Println(Range(1, 10, 2))
	fmt.Println(Range(1, 10, 3))
	fmt.Println(Range(1, 10, 4))
	fmt.Println(Range(1, 10, 11))
	fmt.Println(Range(10, 1, -1))
	fmt.Println(Range(10, 1, -2))
	fmt.Println(Range(10, 1, -3))
	fmt.Println(Range(10, 1, -4))
	fmt.Println(Range(10, 1, -5))
	fmt.Println(Range(10, 1, -6))
	fmt.Println(Range(10, 1, -7))
	fmt.Println(Range(10, 1, -8))
	fmt.Println(Range(10, 1, -9))
	fmt.Println(Range(10, 1, -10))
	fmt.Println(Range(-1, 1, -10))
	fmt.Println(Range(10, 1, -10))
}
