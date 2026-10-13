// Package weather provides the current weather condition in various cities.
package weather

// CurrentCondition stores current weather condition as a string.
var CurrentCondition string

// CurrentLocation stores current weather location as a string.
var CurrentLocation string

// Forecast function takes city and condition as inputs and returns a sentence with CurrentCondition in CurrentLocation.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
