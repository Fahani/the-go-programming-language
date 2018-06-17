// Exercise 4.6: Write an in-place function that squashes each run of adjacent Unicode spaces...
package main

import (
	"fmt"
	"unicode"
)

func squashes(a []byte)[]byte {
	lenA := len(a)
	i := 0
	for ; i<lenA-1; i++{
		if unicode.IsSpace(rune(a[i])) {
			a[i] = ' '
			if unicode.IsSpace(rune(a[i+1])) {
				copy(a[i:], a[i+1:])
				lenA--
				i-- // Stay in the same place to see if the next rune is also a space
			}
		}
	}
	return a[:lenA]
}

func main(){
	s := []byte{'\n',' ','a',' ',' ','d',' ',' ','b'}
	fmt.Printf("%v\n",s)
	fmt.Printf("%v",squashes(s))
}