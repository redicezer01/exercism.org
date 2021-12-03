// package protein provide functions that translates RNA sequences to proteins.
package protein

import "errors"

var ErrStop = errors.New("ErrStop")
var ErrInvalidBase = errors.New("ErrInvalidBase")

// map Codont to Protein.
var CodonsProteinTable = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromRNA translates rna string to slice of proteins
func FromRNA(rna string) ([]string, error) {
	sl := splitStrN(rna, 3)
	r := make([]string, 0, len(sl))
	for _, v := range sl {
		protein, err := FromCodon(v)
		if err == ErrStop {
			return r, nil
		}
		if err != nil {
			return r, err
		}
		r = append(r, protein)
	}
	return r, nil
}

// FromCodon returns protein for providet codon according to CodonsProteinTable
func FromCodon(codon string) (string, error) {
	protein, ok := CodonsProteinTable[codon]
	if !ok {
		return protein, ErrInvalidBase
	}
	if protein == "STOP" {
		return "", ErrStop
	}
	return protein, nil
}

// splitStrN splits string by n rune.
func splitStrN(s string, n int) []string {
	l := len(s)
	sr := []rune(s)
	sl := make([]string, 0, l/n+1)
	chunkSize := 0
	chunk := make([]rune, 0, n)
	for i := 0; i < l; i++ {
		chunk = append(chunk, sr[i])
		chunkSize++
		if chunkSize == n {
			sl = append(sl, string(chunk))
			chunk = make([]rune, 0, n)
			chunkSize = 0
		}
	}
	return sl
}
