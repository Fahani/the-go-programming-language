// Package tempconv performs Celsius and Fahrenheit temperature computations.
package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func main() {
	fmt.Printf("Boiling - Freezing in Celsious %g\n", BoilingC - FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("Boiling in Fahrenheit %g\n", boilingF)
	fmt.Printf("Boiling - Freezing in Fahrenheit %g\n", boilingF-CToF(FreezingC))
	//fmt.Printf("Boiling Fahrenheit - Freezing in Celsious %g\n", boilingF-FreezingC) // (mismatched types Fahrenheit and Celsius)
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"
	//fmt.Println(c == f) // compile error: type mismatch
	fmt.Println(c == Celsius(f)) // "true"!

	fmt.Println(BoilingC)


	fmt.Println(BoilingC.String()) // "100°C"
	fmt.Printf("%v\n", BoilingC) // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", BoilingC) // "100°C"
	fmt.Println(BoilingC) // "100°C"
	fmt.Printf("%g\n", BoilingC) // "100"; does not call String
	fmt.Println(float64(BoilingC)) // "100"; does not call String

}