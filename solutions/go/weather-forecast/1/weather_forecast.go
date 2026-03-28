// Package weather provides tools to display the curent given conditions
// for the given curent location.
package weather

// CurrentCondition string is a variable holding the curent conditions.
var CurrentCondition string

// CurrentLocation string is a variable holding the curent location.
var CurrentLocation string

// Forecast takes two strings city and condition and returns a 
// string with the forecast for the curent lopcation.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
