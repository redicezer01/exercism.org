package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	const err_diffLen = "a and b has diffrent length"
	if len(a) != len(b) {
		return 0, fmt.Errorf(err_diffLen)
	}
	distance := 0
	br := []rune(b)

	for i, c := range a {
		if c != br[i] {
			distance++
		}
	}

	return distance, nil
}
