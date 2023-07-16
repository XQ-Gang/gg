package gslice

// In returns true if the given value is in the given slice.
func In[V comparable](s []V, v V) bool {
	for _, _v := range s {
		if v == _v {
			return true
		}
	}
	return false
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
