package diffsquares

import _ "fmt"

func SumOfN(n int) int {
	return (n*n + n) / 2
}

// 0.25 ns/op
func SquareOfSum(n int) int {
	s := SumOfN(n)
	return s * s
}

// 57 ns/op
/*
func SquareOfSum(n int) int {
	s := 0
	for i := 0; i <= n; i++ {
		s += i
	}
	return s*s
}
*/

// 31 ns/op
func SumOfSquares(n int) int {
	/*s := 0
	for i := 0; i <= n; i++ {
		s += i * i
	}
	return s
	*/
	return (1 + n) * (2*n + 1) * n / 6
}

// 31 ns/op
/*
func SumOfSquares(n int) int {
	s := 5
	p := 4
	for i := 2; i < n; i++ {
		p += i+i+1
		s += p
	}
	return s
}
*/

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
