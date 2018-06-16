// Exercise 4.3: Rewrite reverse to use an array pointer instead of a slice.
package main

import (
	"fmt"
)

func reversePointer(a *[5]int) {
	for i, j := 0, len(a)-1; i<j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	a := [...]int{1,2,3,4,5}
	reversePointer(&a)
	fmt.Println(a)
}