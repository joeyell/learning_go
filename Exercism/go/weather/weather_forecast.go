// Package weather provides a way to print the current location and weather condtions.
package weather

// CurrentCondition represents the current tempreture.
var CurrentCondition string

// CurrentLocation represents the current location.
var CurrentLocation string

// Forecast provides a way to print a string containing the current location and the current weather conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
