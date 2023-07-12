package gptr

func Of[T any](v T) *T {
	return &v
}

func To[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}
