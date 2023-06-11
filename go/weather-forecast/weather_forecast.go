// Package weather provides a tool to tell the forecast
// of a location.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string
// CurrentLocation representes the current locationt that weather forecast should be displayed.
var CurrentLocation string

// Forecast returns a string with information about the the forecast using the CurrentLocation and CurrentCondition values.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
