// Ftoc prints two Fahrenheit to Celsius conversions.
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("Freezing point of water %gºF = %gºC\n", freezingF, fahrenheitToCelsius(freezingF))
	fmt.Printf("Boiling point of water %gºF = %gºC\n", boilingF, fahrenheitToCelsius(boilingF))
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}