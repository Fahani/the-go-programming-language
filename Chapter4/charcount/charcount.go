// Charcount computes counts of Unicode characters.
package main

import (
	"unicode/utf8"
	"os"
	"bufio"
	"io"
	"fmt"
	"unicode"
)

func main() {
	counts := make(map[rune]int) // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0// count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin) // count of
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		//fmt.Fprintf(os.Stdout,"%q\n", r)
	}
	fmt.Println("rune\tcount")
	for k, v := range counts {
		fmt.Printf("%q\t%d\n", k, v)
	}
	fmt.Println("\nlen\tcount")
	for k, v := range utflen {
		if k > 0 {
			fmt.Printf("%d\t%d\n", k, v)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}