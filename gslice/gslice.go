package gslice

import (
	"sort"

	"github.com/XQ-Gang/gg/gutil"
	"github.com/XQ-Gang/gg/gval"

	"github.com/XQ-Gang/gg/constraints"
)

// In returns true if the given value is in the given slice.
func In[V comparable](s []V, v V) bool {
	return Index(s, v) != -1
}

// Index return the index of the given value in the given slice.
// If the given value is not in the given slice, -1 is returned.
func Index[V comparable](s []V, v V) int {
	for i, _v := range s {
		if v == _v {
			return i
		}
	}
	return -1
}

// Filter returns a new slice containing only the elements that pass the given predicate function.
// If the given slice is nil, nil is returned.
func Filter[V any](s []V, f func(V, int) bool) []V {
	if s == nil {
		return nil
	}
	res := make([]V, 0, len(s)/2)
	for i, v := range s {
		if f(v, i) {
			res = append(res, v)
		}
	}
	return res
}

// FilterV is a simplified version of Filter.
// If we don't need the index, we can use FilterV instead of Filter.
func FilterV[V any](s []V, f func(V) bool) []V {
	return Filter(s, func(v V, _ int) bool { return f(v) })
}

// Map returns a new slice containing the results of applying the given function to each element.
// If the given slice is nil, nil is returned.
func Map[V1, V2 any](s []V1, f func(V1, int) V2) []V2 {
	if s == nil {
		return nil
	}
	res := make([]V2, len(s))
	for i, v := range s {
		res[i] = f(v, i)
	}
	return res
}

// MapV is a simplified version of Map.
// If we don't need the index, we can use MapV instead of Map.
func MapV[V1, V2 any](s []V1, f func(V1) V2) []V2 {
	return Map(s, func(v V1, _ int) V2 { return f(v) })
}

// Reduce apply the given function cumulatively to each element of the given slice, from left to right,
// to reduce the slice to a single value. The initial value, if given, is used as the initial value.
// If the initial value is not given, the zero value of the type of the result is used.
// If the given slice is nil, the initial value is returned.
func Reduce[V, T any](s []V, f func(T, V, int) T, initial ...T) T {
	res := gval.Zero[T]()
	if len(initial) > 0 {
		res = initial[0]
	}
	for i, v := range s {
		res = f(res, v, i)
	}
	return res
}

// ReduceV is a simplified version of Reduce.
// If we don't need the index, we can use ReduceV instead of Reduce.
func ReduceV[V, T any](s []V, f func(T, V) T, initial ...T) T {
	return Reduce(s, func(t T, v V, _ int) T { return f(t, v) }, initial...)
}

// ToMap returns a new map containing the results of applying the given function to each element.
// If the given slice is nil, nil is returned.
func ToMap[V1, V2 any, K comparable](s []V1, f func(V1, int) (K, V2)) map[K]V2 {
	if s == nil {
		return nil
	}
	res := make(map[K]V2, len(s))
	for i, v1 := range s {
		k, v2 := f(v1, i)
		res[k] = v2
	}
	return res
}

// All returns true if all elements in the given slice pass the given predicate function.
func All[V any](s []V, f func(V, int) bool) bool {
	for i, v := range s {
		if !f(v, i) {
			return false
		}
	}
	return true
}

// Any returns true if any element in the given slice passes the given predicate function.
func Any[V any](s []V, f func(V, int) bool) bool {
	for i, v := range s {
		if f(v, i) {
			return true
		}
	}
	return false
}

// Sort sorts the given slice in increasing order.
func Sort[V constraints.Ordered](s []V) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func Range[T constraints.Integer](start, end T, step int) []T {
	if step == 0 || start == end || (start < end && step < 0) || (start > end && step > 0) {
		return []T{}
	}
	asc := start < end && step > 0
	stepT := T(step)
	resLen := gutil.If(asc, (end-start+stepT-1)/stepT, gval.Abs((end-start+stepT+1)/stepT))
	res := make([]T, resLen)
	for i := start; gutil.If(asc, i < end, i > end); i += stepT {
		res[(i-start)/stepT] = i
	}
	return res
}
