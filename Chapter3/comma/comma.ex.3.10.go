// Add a , every three characters, non-recursive version. Exercise 3.10
package main

import (
	"fmt"
	"bytes"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	i := 3
	var buf bytes.Buffer
	for ;i < n; i += 3 {
		buf.WriteString(s[i-3:i])
		buf.WriteByte(',')
	}
	if n%i != 0 || n==i {
		buf.WriteString(s[i-3:])
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("123456"))
}
