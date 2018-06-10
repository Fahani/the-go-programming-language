// Returns the basename of the given path/file
// A simpler version using the package strings
package main

import (
	"fmt"
	"strings"
)

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	fmt.Println(basename2("a.go"))
}
