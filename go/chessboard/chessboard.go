package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	r, ok := cb[rank]
	if !ok {
		return 0
	}
	t := 0
	for _, v := range r {
		if v {
			t++
		}
	}
	return t
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	if file < 1 || 8 < file {
		return 0
	}
	t := 0
	for _, v := range cb {
		if v[file-1] {
			t++
		}
	}
	return t
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	t := 0
	for _, _ = range cb {
		t += 8
	}
	return t
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	t := 0
	for _, v := range cb {
		for _, o := range v {
			if o {
				t++
			}
		}
	}
	return t
}
