package gptr

import (
	"testing"

	. "github.com/XQ-Gang/gg/assert"
)

func TestOf(t *testing.T) {
	Eq(t, 1, *Of(1))
	Eq(t, 1, **Of(Of(1)))
	Eq(t, 1, ***Of(Of(Of(1))))
}

func TestTo(t *testing.T) {
	Eq(t, 1, To(Of(1)))
	Eq(t, 1, To(To(Of(Of(1)))))
	Eq(t, 1, To(To(To(Of(Of(Of(1)))))))
	Eq(t, 0, To[int](nil))
}
