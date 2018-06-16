// Given a slice of string, return a slice of string with the non-empty strings of the input
// It may alter the same underlying array
package main

import "fmt"

func nonempty(strings []string)[] string{
	i :=0
	for _,s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main(){
	s := []string{"","a","","b",""}
	fmt.Printf("%q\n",s)
	fmt.Printf("%q\n",nonempty(s))
}