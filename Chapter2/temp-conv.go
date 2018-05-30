// Using the package tempconv
package main

import (
	"fmt"
	"./tempconv"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(tempconv.KToC(tempconv.Kelvin(0)))
}
