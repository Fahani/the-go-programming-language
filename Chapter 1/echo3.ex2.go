// Echo3, exercise 2. Modify the program to print the index and the value, one per line
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i,v := range os.Args[:] {
		fmt.Print("index: ")
		fmt.Print(i)
		fmt.Println(strings.Join(", value: ", v))
	}
}