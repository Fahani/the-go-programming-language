// Echo3, exercise 1. Modify the program to also print the position 0 of the Args array
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[:])
}