// Package weather is an exercise package
// where you have a method and two public variables and
// should add comments to them.
package weather

// CurrentCondition stores the weather condition provided to Forecast method.
var CurrentCondition string

// CurrentLocation stores the location provided to Forecast method.
var CurrentLocation string

// Forecast defines the a weather condition for a city
// and returns a formated text.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
