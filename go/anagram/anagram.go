// package anagram detects anagram words in given list for subject.
package anagram

import "strings"

const abcLen = 26
const normc = 'a'

// Detect returns list of word from candidates that are anagrams to subject word.
func Detect(subject string, candidates []string) []string {
	in := strings.ToLower(subject)
	anagrams := make([]string, 0, len(candidates))
	lmap := getLetterMap(in)

	for _, candidate := range candidates {
		cand := strings.ToLower(candidate)
		clmap := getLetterMap(cand)
		if lmap == clmap && in != cand {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}

// getLetterMap returns letter map for string
// letter map is how much each letter used in word
func getLetterMap(s string) (lmap [abcLen]byte) {
	for _, r := range s {
		if r-normc >= abcLen || r < normc {
			continue
		}
		lmap[r-normc] += 1
	}
	return
}
