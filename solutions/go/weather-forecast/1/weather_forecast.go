// Package weather provides current weather forecasts for cities.
package weather

var (
	// CurrentCondition describes the current weather condition.
	CurrentCondition string
	// CurrentLocation identifies the city for the current forecast.
	CurrentLocation  string
)

// Forecast returns a description of the current weather for the specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
