// Add a , every three numbers, not taking in account the . for floating point numbers and the sign -
package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string{
	var buf bytes.Buffer
	count := 0;
	for _,r := range s {
		if r != '-' && r != '.'{
			count ++
		}
		buf.WriteRune(r)
		if count == 3 {
			buf.WriteByte(',')
			count = 0
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("-124.3456"))
}