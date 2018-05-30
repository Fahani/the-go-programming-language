// Echo4 prints its command-line arguments taking in account the passed flags
package main

import (
	"flag"
	"fmt"
	"strings"
	"os"
)

var n = flag.Bool("n", false, "omit the trailing newline")
var sep = flag.String("sep", " ", "separator")

func main() {
	fmt.Println(strings.Join(os.Args[1:], "-"))
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
