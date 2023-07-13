package gmap

// Keys returns a slice of keys of the given map.
// The order of the keys is not guaranteed.
// If the given map is nil, nil is returned.
func Keys[K comparable, V any](m map[K]V) []K {
	if m == nil {
		return nil
	}
	res := make([]K, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}

// Values returns a slice of values of the given map.
// The order of the values is not guaranteed.
// If the given map is nil, nil is returned.
func Values[K comparable, V any](m map[K]V) []V {
	if m == nil {
		return nil
	}
	res := make([]V, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// Get returns the value of the given key in the given map.
// If the key does not exist, the default value is returned.
func Get[K comparable, V any](m map[K]V, k K, def V) V {
	if v, ok := m[k]; ok {
		return v
	}
	return def
}

// Filter returns a new map containing only the key-value pairs that pass the given predicate function.
// If the given map is nil, nil is returned.
func Filter[K comparable, V any](m map[K]V, f func(K, V) bool) map[K]V {
	if m == nil {
		return nil
	}
	res := make(map[K]V, len(m)/2)
	for k, v := range m {
		if f(k, v) {
			res[k] = v
		}
	}
	return res
}

// Map returns a new map containing the results of applying the given function to each key-value pair.
// If the given map is nil, nil is returned.
func Map[K1, K2 comparable, V1, V2 any](m map[K1]V1, f func(K1, V1) (K2, V2)) map[K2]V2 {
	if m == nil {
		return nil
	}
	res := make(map[K2]V2, len(m))
	for k1, v1 := range m {
		k2, v2 := f(k1, v1)
		res[k2] = v2
	}
	return res
}
