package gptr

// Of returns a pointer to the given value.
func Of[T any](v T) *T {
	return &v
}

// To returns the value pointed to by the given pointer.
// If the pointer is nil, the zero value of the type is returned.
func To[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}
