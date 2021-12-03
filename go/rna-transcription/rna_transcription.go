// package strand provide function to convert DNA sequance to RNA
package strand

import "bytes"

// transcription table
var transcribeTab = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA takes dna sequance and converts rna according to transcribeTab
func ToRNA(dna string) string {
	var rna bytes.Buffer

	for _, r := range dna {
		nuc := r
		if replace, ok := transcribeTab[r]; ok {
			nuc = replace
		}
		rna.WriteRune(nuc)
	}
	return rna.String()
}
