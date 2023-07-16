package gmap

import (
	"testing"

	. "github.com/XQ-Gang/gg/assert"
	"github.com/XQ-Gang/gg/gslice"
)

func TestKeys(t *testing.T) {
	type args struct {
		m map[int]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"nil", args{nil}, nil},
		{"empty", args{map[int]int{}}, []int{}},
		{"normal", args{map[int]int{1: 2, 3: 4, 5: 6}}, []int{1, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Keys(tt.args.m)
			gslice.Sort(actual)
			Eq(t, tt.want, actual)
		})
	}
}

func TestValues(t *testing.T) {
	type args struct {
		m map[int]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"nil", args{nil}, nil},
		{"empty", args{map[int]int{}}, []int{}},
		{"normal", args{map[int]int{1: 2, 3: 4, 5: 6}}, []int{2, 4, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Values(tt.args.m)
			gslice.Sort(actual)
			Eq(t, tt.want, actual)
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		m   map[int]int
		k   int
		def int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"nil", args{nil, 1, 9}, 9},
		{"empty", args{map[int]int{}, 1, 9}, 9},
		{"exist", args{map[int]int{1: 2, 3: 4, 5: 6}, 1, 9}, 2},
		{"not_exist", args{map[int]int{1: 2, 3: 4, 5: 6}, 2, 9}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eq(t, tt.want, Get(tt.args.m, tt.args.k, tt.args.def))
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		m map[int]int
		f func(int, int) bool
	}
	filterFunc := func(k int, v int) bool { return k == 1 || v == 1 }
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{"nil", args{nil, filterFunc}, nil},
		{"empty", args{map[int]int{}, filterFunc}, map[int]int{}},
		{"normal", args{map[int]int{1: 2, 3: 4, 5: 6}, filterFunc}, map[int]int{1: 2}},
		{"all", args{map[int]int{1: 2, 3: 4, 5: 6}, func(_ int, _ int) bool { return true }}, map[int]int{1: 2, 3: 4, 5: 6}},
		{"none", args{map[int]int{1: 2, 3: 4, 5: 6}, func(_ int, _ int) bool { return false }}, map[int]int{}},
		{"only_key", args{map[int]int{1: 2, 3: 4, 5: 6}, func(k int, _ int) bool { return k > 3 }}, map[int]int{5: 6}},
		{"only_value", args{map[int]int{1: 2, 3: 4, 5: 6}, func(_ int, v int) bool { return v < 4 }}, map[int]int{1: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eq(t, tt.want, Filter(tt.args.m, tt.args.f))
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		m map[int]int
		f func(int, int) (int, int)
	}
	mapFunc := func(k int, v int) (int, int) { return k + 1, v + 1 }
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{"nil", args{nil, mapFunc}, nil},
		{"empty", args{map[int]int{}, mapFunc}, map[int]int{}},
		{"normal", args{map[int]int{1: 2, 3: 4, 5: 6}, mapFunc}, map[int]int{2: 3, 4: 5, 6: 7}},
		{"only_key", args{map[int]int{1: 2, 3: 4, 5: 6}, func(k int, _ int) (int, int) { return k + 1, 0 }}, map[int]int{2: 0, 4: 0, 6: 0}},
		{"only_value", args{map[int]int{1: 2, 3: 4, 5: 6}, func(_ int, v int) (int, int) { return v + 1, 0 }}, map[int]int{3: 0, 5: 0, 7: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eq(t, tt.want, Map(tt.args.m, tt.args.f))
		})
	}
}
