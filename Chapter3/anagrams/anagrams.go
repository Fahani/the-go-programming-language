// Exercise 3.12. Write a function that reports if two strings are anagram
package main

import "fmt"

func areAnagram(a, b string) bool {
	runesA, runesB := []rune(a), []rune(b)
	lenA, lenB := len(runesA), len(runesB)

	if lenA != lenB {
		return false
	}

	for _, rA := range runesA {
		for i, rB := range runesB {
			 if rA == rB {
			 	runesB = append(runesB[:i],runesB[i+1:]...)
			 	lenB = len(runesB)
			 	break
			 }

			 if lenB == i+1 {
			 	return false
			 }
		}
	}
	return true
}

func main() {
	fmt.Println(areAnagram("ANAGRAM ", "GARANAMA"))
}