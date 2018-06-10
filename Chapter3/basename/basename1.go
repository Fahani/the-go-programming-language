// Returns the basename of the given path/file
package main

import "fmt"

func basename(s string) string {
	for i := len(s)-1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s)-1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	fmt.Println(basename("/aasd7/7777/a.c.go"))
}
