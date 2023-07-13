package gstr

type A string

// Len returns the number of runes in the given string.
// This is equivalent to utf8.RuneCountInString.
func Len[S ~string](s S) int {
	rs := []rune(s)
	return len(rs)
}
