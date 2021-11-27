// package robotname manage robot factory settings
package robotname

import "math/rand"
import "fmt"
import "errors"

// total number of unique robot name in format: {LLDDD},
//	L - capital latin letter,
//	D - digit (0..9)
const numOfComb = 26 * 26 * 10 * 10 * 10

// usedNames map of all generated names
var usedNames = map[string]bool{}

// Define the Robot type here.
type Robot struct {
	name string
}

// Name return robots name or generates it
func (r *Robot) Name() (string, error) {
	// try generate name.
	if r.name == "" {
		r.Reset()
	}
	// if can't generate then names exhausted.
	if r.name == "" {
		return "", errors.New("Run out of robot names.")
	}

	return r.name, nil
}

// Reset try to generate new uniqe name for robot or return "" if names exhausted.
func (r *Robot) Reset() {
	notDone := true
	// try generate name
	for notDone {
		r.name = genName(r.name)
		if !usedNames[r.name] {
			usedNames[r.name] = true
			notDone = false
		} else if len(usedNames) == numOfComb {
			notDone = false
			r.name = ""
		}

	}
}

// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
// BenchmarkName-8          1184066              1215 ns/op
// Banchmark for Name()

// genName generates random name in format AA123.
func genName(name string) string {
	var l1, l2, d1, d2, d3 int32
	// if empty generate whole name
	if name == "" {
		l1, l2 = rand.Int31n(26), rand.Int31n(26)
		d1, d2, d3 = rand.Int31n(10), rand.Int31n(10), rand.Int31n(10)
	} else { // if try generate one symbol to get new unique name
		nr := []rune(name)
		r := rand.Int31n(5) // randomly choose which symbol to generate
		chooseFrom := int32(26)
		add := int32('A')
		if r > 1 {
			add = int32('0')
			chooseFrom = 10
		}
		nr[r] = rand.Int31n(chooseFrom) + add
		l1, l2, d1, d2, d3 = nr[0]-'A', nr[1]-'A', nr[2]-'0', nr[3]-'0', nr[4]-'0'
	}
	return fmt.Sprintf("%c%c%d%d%d", l1+'A', l2+'A', d1, d2, d3)
}

// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
// BenchmarkName-8          1000000              4620 ns/op
// Banchmark for Name()

/*
func genName() string {
	l1 := rand.Int31n(26)
	l2 := rand.Int31n(26)
	d1 := rand.Int31n(10)
	d2 := rand.Int31n(10)
	d3 := rand.Int31n(10)
	return fmt.Sprintf("%c%c%d%d%d", l1+'A', l2+'A', d1, d2, d3)
}
*/
