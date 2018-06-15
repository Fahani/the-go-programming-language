// Example append a int to a int slice, making it grow if need it
package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a builtin function; see text
	}
	z[len(x)] = y
	return z
}

func main() {
	x := make([]int, 4, 4)
	y := 4
	fmt.Println(appendInt(x,y))
}