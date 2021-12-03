// package romannumerals represents numbers in roman numerals
package romannumerals

import "math"
import "strings"
import "strconv"
import "fmt"

// Roman digits
var romanDigits = map[int32]rune{
	1:     'I',
	5:     'V',
	10:    'X',
	50:    'L',
	100:   'C',
	500:   'D',
	1000:  'M',
	5000:  '_',
	10000: '_',
}

// Repeated pattern for building roman numbers
// to avoid collision using L, M, H with roman digits:
// 	1 - means Low
// 	5 - means Mid
// 	9 - means High
var patterns = map[int32]string{
	0:  "",
	1:  "1",
	2:  "11",
	3:  "111",
	4:  "15",
	5:  "5",
	6:  "51",
	7:  "511",
	8:  "5111",
	9:  "19",
	10: "9",
}

// Low, Mid, Hi is roman digits that need to build number according to pattern
type LowMidHi struct {
	L rune
	M rune
	H rune
}

// getRomanDigitByRank returns roman digit (Low, Mid or Hi i int32)  according to position of digit
func getRomanDigitByRank(rank, i int32) rune {
	pos := float64(i) * math.Pow(10, float64(rank))
	return romanDigits[int32(pos)]
}

// getLowMidHi returns LowMidHi type
func getLowMidHi(rank int32) (lmh LowMidHi) {
	lmh.L = getRomanDigitByRank(rank, 1)
	lmh.M = getRomanDigitByRank(rank, 5)
	lmh.H = getRomanDigitByRank(rank, 10)
	return
}

// ToRomanNumeral converts regular number to Roman Numeral
func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3000 {
		return "", fmt.Errorf("invalid input %v, should be 0 < input < 3000.", input)
	}
	var roman strings.Builder
	inStr := strconv.Itoa(input)
	inStrLen := len(inStr)
	for i, r := range inStr {
		if r-'0' == 0 {
			continue
		}
		lmh := getLowMidHi(int32(inStrLen - 1 - i))
		pattern := patterns[r-'0']
		pattern = strings.Replace(pattern, "1", string(lmh.L), -1)
		pattern = strings.Replace(pattern, "5", string(lmh.M), -1)
		pattern = strings.Replace(pattern, "9", string(lmh.H), -1)
		_, _ = roman.WriteString(pattern)
	}
	return roman.String(), nil
}
