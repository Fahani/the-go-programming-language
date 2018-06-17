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
			i-- // stay in the same place to see of the next adjacent is the same
		}
	}
	return a[:lenA]
}

func main(){
	s := []string{"b","b"}
	fmt.Println(removeAdDu(s))
}