package wordcount

import "strings"
import "fmt"
 
type Frequency map[string]int

func WordCount(phrase string) Frequency {
	res := Frequency{}

	fmt.Println(splitByWhiteSpace(phrase))
	return res
}

func splitByWhiteSpace(s string) []string {
	res := []string{}
	l := len([]rune(s))
	word := strings.Builder{}
	for i, r := range s {
		isWhiteSpace := false
		if r == ' ' || r == '\n' || r == '\t' {
			isWhiteSpace = true
		}
		if !isWhiteSpace {
			word.WriteRune(r)
		}
		if isWhiteSpace || i == l-1 {
			if word.Len() != 0 {
				res = append(res, word.String())
				word = strings.Builder{}
			}
		}
			
	}
	return res
}
			
	
