// Add a , every three characters
package main

import "fmt"

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return s[:3] + "," + comma(s[3:])
}

func main() {
	fmt.Println(comma("123456"))
}
