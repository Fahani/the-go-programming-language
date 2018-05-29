// This file contains the functions to convert between the different units
package unitconv

// CToF converts a Celsius temperature to Fahrenheit
func CToF (c Celsius) Fahrenheit {
	return Fahrenheit(c * 9 / 5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius
func FToC (f Fahrenheit) Celsius{
	return Celsius((f - 32) * 5 / 9)
}

// MToF converts a Meter unit to Feet
func MToF (m Meter) Feet {
	return Feet(m * 3.28084)
}

// FToM Converts a Feet unit to Meter
func FToM (f Feet) Meter {
	return Meter(f / 3.28084)
}

// KToP converts a Kilogram unit to Pound
func KToP (k Kilogram) Pound {
	return Pound(k * 2.20462)
}

// PToK converts a Pound unit to Kilogram
func PToK (p Pound) Kilogram {
	return Kilogram(p / 2.20462)
}