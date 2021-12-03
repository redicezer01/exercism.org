// package triangle provide function that determines triangle kind
package triangle

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides return triangle kind in provided sides
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	zeroSide := a <= 0 || b <= 0 || c <= 0
	sumTwoSidesSmaller := a+b < c || a+c < b || b+c < a

	if zeroSide || sumTwoSidesSmaller {
		k = NaT
	} else if a == b && b == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}
	return k
}
