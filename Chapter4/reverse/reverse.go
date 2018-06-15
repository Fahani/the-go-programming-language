package main

import "fmt"

func reverse(a []int) []int {
	for i, j := 0, len(a)-1; i<j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func main() {
	a := [...]int{1, 2, 3, 4}
	fmt.Println(reverse(a[:]))
}