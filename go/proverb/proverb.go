// package proverb generates proverbial rhyme for a given input
package proverb

// Proverb returns generated proverb in a slice
func Proverb(rhyme []string) []string {
	res := make([]string, 0, len(rhyme))
	if rhyme == nil || len(rhyme) == 0 {
		return res
	}
	for i := 0; i < len(rhyme) - 1; i++ {
		res = append(res, "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost.")
	}	
	res = append(res, "And all for the want of a " + rhyme[0] + ".")
	return res
}
