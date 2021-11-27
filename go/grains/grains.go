package grains

import "errors"

// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
// BenchmarkSquare-8       121521229                9.884 ns/op

func Square(number int) (uint64, error) {
	number -= 1
	if number < 0 || number > 63 {
		return 0, errors.New("Out of range 64.")
	}
	return uint64(1) << number, nil
}

// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
// BenchmarkTotal-8        65443152                18.06 ns/op

func Total() uint64 {
	var s uint64
	for i := 0; i <= 64; i++ {
		s |= 1
		s <<= 1
	}
	s |= 1
	return s
}
