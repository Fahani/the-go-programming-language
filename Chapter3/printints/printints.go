// Using the type bytes.Buffer to write strings/int/runes/ in it
package main

import (
	"bytes"
	"fmt"
)

func printInts(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(printInts([]int{1,2,3}))
}