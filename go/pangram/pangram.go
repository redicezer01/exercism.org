// package pangram determine if sentence is pangram
package pangram

const (
	a    = byte('a')
	z    = byte('z')
	A    = byte('A')
	Z    = byte('Z')
	dist = a - A
)

// IsPangram returns true if sentence is pangram
func IsPangram(input string) bool {
	cnt := 0             // how much letters from ABC was met
	counted := uint32(0) // each bit set means that letter was in sentence
	for _, r := range input {
		letter := toLower(byte(r))
		// count only unique letters
		if !letterCounted(letter, counted) && isLetter(letter) {
			counted |= (uint32(1) << uint32(letter-a))
			cnt++
		}
		// if count of unique letters >= to 26 (number of letters in ABC)
		if cnt >= 26 {
			return true
		}
	}
	return false
}

// toLower custom func converts ASCII byte to lower code
func toLower(letter byte) byte {
	if A <= letter && letter <= Z {
		return letter + dist
	}
	return letter
}

// isLetter custom func determines whether ASCII byte is letter or not
func isLetter(letter byte) bool {
	if a <= toLower(letter) && toLower(letter) <= z {
		return true
	}
	return false
}

// letterCounted func helper for IsPangram that counts unique ASCII byte symbols
func letterCounted(letter byte, i uint32) bool {
	if i&(uint32(1)<<uint32(letter-a)) != 0 {
		return true
	}
	return false
}
