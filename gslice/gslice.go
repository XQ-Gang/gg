package gslice

func In[V comparable](s []V, v V) bool {
	for _, _v := range s {
		if v == _v {
			return true
		}
	}
	return false
}

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

func Map[V1, V2 any](s []V1, f func(V1, int) V2) []V2 {
	if s == nil {
		return nil
	}
	res := make([]V2, 0, len(s))
	for i, v := range s {
		res[i] = f(v, i)
	}
	return res
}

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

func All[V any](s []V, f func(V, int) bool) bool {
	for i, v := range s {
		if !f(v, i) {
			return false
		}
	}
	return true
}

func Any[V any](s []V, f func(V, int) bool) bool {
	for i, v := range s {
		if f(v, i) {
			return true
		}
	}
	return false
}
