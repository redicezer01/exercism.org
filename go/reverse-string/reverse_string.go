// package reverse reverse string.
package reverse

import "strings"

// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
// BenchmarkReverse-8       2008104               642.2 ns/op
//
// Reverse returns reversed string.
func Reverse(input string) string {
	rs := []rune(input)
	l := len(rs)
	sb := strings.Builder{}
	for i := l - 1; i >= 0; i-- {
		sb.WriteRune(rs[i])
	}
	return sb.String()
}

// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
// BenchmarkReverse-8        371457              5538 ns/op
//
// recursive method.
// Reverse returns reversed string.
/*
func Reverse(in string) string {
	sr := []rune(in)
	l := len(sr)
	if l == 0 || l == 1 {
		return in
	}
	return Reverse(string(sr[1:])) + string(sr[0])
}
*/
