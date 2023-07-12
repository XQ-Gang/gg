package gstr

type A string

func Len[S ~string](s S) int {
	rs := []rune(s)
	return len(rs)
}
