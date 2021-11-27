package raindrops

import "fmt"

func Convert(number int) (r string) {
	// n := []rune(fmt.Sprintf("%d",number))
	if number%3 == 0 {
		r += "Pling"
	}
	// if (n[len(n)-1]-'0') == 5 ||  (n[len(n)-1]-'0') == 0 {
	if number%5 == 0 {
		r += "Plang"
	}
	if number%7 == 0 {
		r += "Plong"
	}

	if r == "" {
		return fmt.Sprintf("%d", number)
	}
	return r
}
