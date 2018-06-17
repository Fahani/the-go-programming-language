// Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.
package main

import "fmt"

func removeAdDu(a []string)[]string {
	lenA := len(a)
	i := 0
	for ; i<lenA-1; i++ {
		if a[i] == a[i+1] {
			copy(a[i:],a[i+1:])
			lenA--
		}
	}
	return a[:len(a)+1-i]
}

func main(){
	s := []string{"a","b", "b", "c", "c", "d"}
	fmt.Println(removeAdDu(s))
}