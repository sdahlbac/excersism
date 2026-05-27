// Package weather provides functionality for forecasting weather
// for a specific city and condition.
package weather

var (
	// CurrentCondition represents the current condition.
	CurrentCondition string
	// CurrentLocation represents the current location.
	CurrentLocation string
)

// Forecast returns the current weather for a specific location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
