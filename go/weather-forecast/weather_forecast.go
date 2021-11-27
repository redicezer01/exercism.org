// Package weather implements function that formats weather condition in location.
package weather

// CurrentCondition stores current weather condition.
var CurrentCondition string 

// CurrentLocation stores current city.
var CurrentLocation string  

// Forecast returns formated string of weather condition in city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
