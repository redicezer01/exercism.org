// package space implements Age function
package space

type Planet string

const secInYear = 31557600

// orbital peroid in Earth years
var planets = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age returns age on specified planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / secInYear / planets[planet]
}
