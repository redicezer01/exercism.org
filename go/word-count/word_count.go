// package wordcount counts frequency of each word in sentence.
package wordcount

import "strings"
import "unicode"

// Frequency type contains word and frequency in sentence.
type Frequency map[string]int

// WordCount returns frequency map for each word in string.
func WordCount(phrase string) Frequency {
	res := Frequency{}

	for _, w := range splitByWhiteSpace(phrase) {
		res[w] += 1
	}

	return res
}

// splitByWhiteSpace returns slice of words parsed from given string.
func splitByWhiteSpace(s string) []string {
	res := []string{}
	l := len([]rune(s))
	word := strings.Builder{}
	should := 0 // 1 - letter, 2 - digit
	prevSymbol := ' '

	for i, r := range s {
		if should == 0 {
			should = shouldBe(r)
		}
		if !isWhiteSpace(r) &&
			((should == 1 && isLetter(r)) ||
				(should == 2 && unicode.IsDigit(r))) {
			word.WriteRune(unicode.ToLower(r))
		}
		if (isWhiteSpace(r) || i == l-1) && word.Len() != 0 {
			if r == '\'' || prevSymbol == '\'' {
				res = append(res, word.String()[:word.Len()-1])
			} else {
				res = append(res, word.String())
			}
			word = strings.Builder{}
			should = 0
		}
		prevSymbol = r
	}

	return res
}

// isPunct exclude ' e.g. can't, don't and so on.
func isPunct(r rune) bool {
	return unicode.IsPunct(r) && r != '\''
}

// isWhiteSpace .
func isWhiteSpace(r rune) bool {
	return unicode.IsSpace(r) || isPunct(r)
}

// isLetter add ' as appropriate word letter.
func isLetter(r rune) bool {
	return unicode.IsLetter(r) || r == '\''
}

// shouldBe returns what kind of word we expect letter word or number.
func shouldBe(r rune) int {
	switch {
	case unicode.IsLetter(r):
		return 1
	case unicode.IsDigit(r):
		return 2
	default:
		return 0
	}
}
