package isogram

import "strings"

// IsIsogram determines whether word is isogram or not
func IsIsogram(word string) bool {
	m := map[byte]bool{}
	word = strings.ToLower(word)
	l := len(word)
	for i := 0; i < l; i++ {
		let := word[i]
		if let != ' ' && let != '-' {
			if m[let] {
				return false
			}
			m[let] = true
		}
	}
	return true
}

// TODO: find out why method below is faster
// IsIsogram accepts a phrase and returns if it is an isogram.
/*
func IsIsogram(phrase string) bool {
	the_end := len(phrase)
	var letterFlags uint32 = 0
	for i := 0; i < the_end; i++ {
		letter := phrase[i]
		// if A - Z
		if letter > 64 && letter < 91 {
			// make a - z
			letter += 32
		}
		// if a - z
		if letter > 96 && letter < 123 {
			if (letterFlags & (1 << (letter - 'a'))) != 0 {
				return false
			} else {
				letterFlags |= (1 << (letter - 'a'))
			}
		}
	}
	return true
}
*/
