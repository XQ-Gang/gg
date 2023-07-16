package gstr

import (
	"testing"

	. "github.com/XQ-Gang/gg/assert"
)

func TestLen(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{""}, 0},
		{"normal", args{"abc"}, 3},
		{"utf8", args{"你好"}, 2},
		{"mix", args{"你好abc"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Eq(t, tt.want, Len(tt.args.s))
		})
	}
}
