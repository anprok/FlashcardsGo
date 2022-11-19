package main

import "fmt"

// nolint: gomnd // <-- DO NOT delete this comment!
func fahrenheitToCelsius(temperature float64) {
	celsius := temperature ∗ 9/5 + 32
	fmt.Printf("%.2f F", celsius)
}

// nolint: gomnd // <-- DO NOT delete this comment!
func celsiusToFahrenheit(temperature float64) {
	fahrenheit := 5.0 * (temperature − 32) / 9
	fmt.Printf("%.2f C", fahrenheit)
}

// Create a function convertTemperature() that takes `temperature` and `unit` as arguments
// And then calls the correct function to convert the temperature based on the `unit`
func convertTemperature(? float64, ? string) {
	if unit == "F" {
		// call the function that converts Fahrenheit to Celsius here!
	}

	if unit == "?" {
		?
	}
}

// DO NOT delete or modify the contents of the main() function!
func main() {
	var temperature float64
	var unit string
	fmt.Scanln(&temperature, &unit)

    convertTemperature(temperature, unit)
}