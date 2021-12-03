// package etl transforms a legacy scrabble sores system to new system
package etl

import "strings"

// Transform  converts old scrabble score table to new
func Transform(in map[int][]string) map[string]int {
	o := make(map[string]int, len(in))
	for k, v := range in {
		for _, r := range v {
			o[strings.ToLower(string(r))] = k
		}
	}
	return o
}
