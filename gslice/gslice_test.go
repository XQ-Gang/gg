package gslice

import (
	"strconv"
	"testing"

	. "github.com/XQ-Gang/gg/assert"
)

func TestIn(t *testing.T) {
	type args struct {
		s []int
		v int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"nil", args{nil, 1}, false},
		{"found", args{[]int{1, 2, 3}, 2}, true},
		{"not found", args{[]int{1, 2, 3}, 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eq(t, tt.want, In(tt.args.s, tt.args.v))
		})
	}
}

func TestIndex(t *testing.T) {
	type args struct {
		s []int
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"nil", args{nil, 1}, -1},
		{"found", args{[]int{1, 2, 3}, 2}, 1},
		{"not found", args{[]int{1, 2, 3}, 4}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eq(t, tt.want, Index(tt.args.s, tt.args.v))
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		s []int
		f func(int, int) bool
	}
	filterFunc := func(v int, i int) bool { return v == 1 || i == 1 }
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"nil", args{nil, filterFunc}, nil},
		{"empty", args{[]int{}, filterFunc}, []int{}},
		{"normal", args{[]int{1, 2, 3}, filterFunc}, []int{1, 2}},
		{"all", args{[]int{1, 2, 3}, func(_ int, _ int) bool { return true }}, []int{1, 2, 3}},
		{"none", args{[]int{1, 2, 3}, func(_ int, _ int) bool { return false }}, []int{}},
		{"only_value", args{[]int{1, 2, 3}, func(v int, _ int) bool { return v%2 == 0 }}, []int{2}},
		{"only_index", args{[]int{1, 2, 3}, func(_ int, i int) bool { return i%2 == 0 }}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eq(t, tt.want, Filter(tt.args.s, tt.args.f))
		})
	}
}

func TestFilterV(t *testing.T) {
	type args struct {
		s []int
		f func(int) bool
	}
	filterFunc := func(v int) bool { return v == 1 }
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"nil", args{nil, filterFunc}, nil},
		{"empty", args{[]int{}, filterFunc}, []int{}},
		{"normal", args{[]int{1, 2, 3}, filterFunc}, []int{1}},
		{"all", args{[]int{1, 2, 3}, func(_ int) bool { return true }}, []int{1, 2, 3}},
		{"none", args{[]int{1, 2, 3}, func(_ int) bool { return false }}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eq(t, tt.want, FilterV(tt.args.s, tt.args.f))
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		s []int
		f func(int, int) int
	}
	mapFunc := func(v int, i int) int { return v + i }
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"nil", args{nil, mapFunc}, nil},
		{"empty", args{[]int{}, mapFunc}, []int{}},
		{"normal", args{[]int{1, 2, 3}, mapFunc}, []int{1, 3, 5}},
		{"only_value", args{[]int{1, 2, 3}, func(v int, _ int) int { return v * 2 }}, []int{2, 4, 6}},
		{"only_index", args{[]int{1, 2, 3}, func(_ int, i int) int { return i * 2 }}, []int{0, 2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eq(t, tt.want, Map(tt.args.s, tt.args.f))
		})
	}
}

func TestMapV(t *testing.T) {
	type args struct {
		s []int
		f func(int) int
	}
	mapFunc := func(v int) int { return v + 1 }
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"nil", args{nil, mapFunc}, nil},
		{"empty", args{[]int{}, mapFunc}, []int{}},
		{"normal", args{[]int{1, 2, 3}, mapFunc}, []int{2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eq(t, tt.want, MapV(tt.args.s, tt.args.f))
		})
	}
}

func TestReduce(t *testing.T) {
	type args struct {
		s       []int
		f       func(int, int, int) int
		initial int
	}
	reduceFunc := func(agg int, val int, idx int) int { return agg + val + idx }
	tests := []struct {
		name string
		args args
		want int
	}{
		{"nil", args{nil, reduceFunc, 0}, 0},
		{"empty", args{[]int{}, reduceFunc, 0}, 0},
		{"initial", args{[]int{}, reduceFunc, 1}, 1},
		{"normal#initial0", args{[]int{1, 2, 3}, reduceFunc, 0}, 9},
		{"normal#initial1", args{[]int{1, 2, 3}, reduceFunc, 1}, 10},
		{"only_value", args{[]int{1, 2, 3}, func(agg int, val int, _ int) int { return agg + val }, 0}, 6},
		{"only_index", args{[]int{1, 2, 3}, func(agg int, _ int, idx int) int { return agg + idx }, 0}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eq(t, tt.want, Reduce(tt.args.s, tt.args.f, tt.args.initial))
		})
	}
}

func TestReduceV(t *testing.T) {
	type args struct {
		s       []int
		f       func(int, int) int
		initial int
	}
	reduceFunc := func(agg int, val int) int { return agg + val }
	tests := []struct {
		name string
		args args
		want int
	}{
		{"nil", args{nil, reduceFunc, 0}, 0},
		{"empty", args{[]int{}, reduceFunc, 0}, 0},
		{"initial", args{[]int{}, reduceFunc, 1}, 1},
		{"normal#initial0", args{[]int{1, 2, 3}, reduceFunc, 0}, 6},
		{"normal#initial1", args{[]int{1, 2, 3}, reduceFunc, 1}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eq(t, tt.want, ReduceV(tt.args.s, tt.args.f, tt.args.initial))
		})
	}
}

func TestFilterMap(t *testing.T) {
	type args struct {
		s []int
		f func(int, int) (int, bool)
	}
	filterMapFunc := func(v int, i int) (int, bool) { return v + i, v%2 == 0 }
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"nil", args{nil, filterMapFunc}, nil},
		{"empty", args{[]int{}, filterMapFunc}, []int{}},
		{"normal", args{[]int{1, 2, 3, 4, 5}, filterMapFunc}, []int{3, 7}},
		{"only_value", args{[]int{1, 2, 3}, func(v int, _ int) (int, bool) { return v * 2, v%2 == 0 }}, []int{4}},
		{"only_index", args{[]int{1, 2, 3}, func(_ int, i int) (int, bool) { return i * 2, i%2 == 0 }}, []int{0, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eq(t, tt.want, FilterMap(tt.args.s, tt.args.f))
		})
	}
}

func TestFilterMapV(t *testing.T) {
	type args struct {
		s []int
		f func(int) (int, bool)
	}
	filterMapFunc := func(v int) (int, bool) { return v + 1, v%2 == 0 }
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"nil", args{nil, filterMapFunc}, nil},
		{"empty", args{[]int{}, filterMapFunc}, []int{}},
		{"normal", args{[]int{1, 2, 3, 4, 5}, filterMapFunc}, []int{3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eq(t, tt.want, FilterMapV(tt.args.s, tt.args.f))
		})
	}
}

func TestToMap(t *testing.T) {
	type args struct {
		s []int
		f func(int, int) (string, int)
	}
	toMapFunc := func(v int, i int) (string, int) { return strconv.Itoa(v), v + i }
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{"nil", args{nil, toMapFunc}, nil},
		{"empty", args{[]int{}, toMapFunc}, map[string]int{}},
		{"normal", args{[]int{1, 2, 3}, toMapFunc}, map[string]int{"1": 1, "2": 3, "3": 5}},
		{"conflict", args{[]int{1, 2, 3}, func(v int, _ int) (string, int) { return "1", v }}, map[string]int{"1": 3}},
		{"only_value", args{[]int{1, 2, 3}, func(v int, _ int) (string, int) { return strconv.Itoa(v), v * 2 }}, map[string]int{"1": 2, "2": 4, "3": 6}},
		{"only_index", args{[]int{1, 2, 3}, func(_ int, i int) (string, int) { return strconv.Itoa(i), i * 2 }}, map[string]int{"0": 0, "1": 2, "2": 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eq(t, tt.want, ToMap(tt.args.s, tt.args.f))
		})
	}
}

func TestAll(t *testing.T) {
	type args struct {
		s []int
		f func(int, int) bool
	}
	allFunc := func(v int, i int) bool { return v > 0 }
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"nil", args{nil, allFunc}, true},
		{"empty", args{[]int{}, allFunc}, true},
		{"normal#true", args{[]int{1, 2, 3}, allFunc}, true},
		{"normal#false", args{[]int{-1, 0, 1}, allFunc}, false},
		{"only_value", args{[]int{1, 2, 3}, func(v int, _ int) bool { return v < 3 }}, false},
		{"only_index", args{[]int{1, 2, 3}, func(_ int, i int) bool { return i < 3 }}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eq(t, tt.want, All(tt.args.s, tt.args.f))
		})
	}
}

func TestAny(t *testing.T) {
	type args struct {
		s []int
		f func(int, int) bool
	}
	anyFunc := func(v int, i int) bool { return v > 0 }
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"nil", args{nil, anyFunc}, false},
		{"empty", args{[]int{}, anyFunc}, false},
		{"normal#true", args{[]int{1, 2, 3}, anyFunc}, true},
		{"normal#false", args{[]int{-2, -1, 0}, anyFunc}, false},
		{"only_value", args{[]int{1, 2, 3}, func(v int, _ int) bool { return v < 1 }}, false},
		{"only_index", args{[]int{1, 2, 3}, func(_ int, i int) bool { return i < 1 }}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eq(t, tt.want, Any(tt.args.s, tt.args.f))
		})
	}
}

func TestSort(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"nil", args{nil}, nil},
		{"empty", args{[]int{}}, []int{}},
		{"normal#unordered", args{[]int{2, 1, 3}}, []int{1, 2, 3}},
		{"normal#asc", args{[]int{3, 2, 1}}, []int{1, 2, 3}},
		{"normal#desc", args{[]int{3, 2, 1}}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.args.s)
			Eq(t, tt.want, tt.args.s)
		})
	}
}

func TestRange(t *testing.T) {
	type args struct {
		start int
		end   int
		step  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"step=0", args{1, 4, 0}, []int{}},
		{"start=end", args{1, 1, 1}, []int{}},
		{"start<end&step<0", args{1, 4, -1}, []int{}},
		{"start>end&step>0", args{4, 1, 1}, []int{}},
		{"start<end&step=1", args{1, 4, 1}, []int{1, 2, 3}},
		{"start<end&step=2", args{1, 4, 2}, []int{1, 3}},
		{"start<end&step>(end-start)", args{1, 4, 5}, []int{1}},
		{"start>end&step=-1", args{4, 1, -1}, []int{4, 3, 2}},
		{"start>end&step=-2", args{4, 1, -2}, []int{4, 2}},
		{"start>end&step<(start-end)", args{4, 1, -5}, []int{4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eq(t, tt.want, Range(tt.args.start, tt.args.end, tt.args.step))
		})
	}
}
