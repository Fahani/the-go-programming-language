// Ex 4.9 counts the frequently of words in a input text file
package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	file, err := os.Open("Chapter4/wordfreq/data.txt")
	words := map[string]int{}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %v", err)
		os.Exit(1)
	}

	read := bufio.NewScanner(file)
	read.Split(bufio.ScanWords)

	for read.Scan(){
		words[read.Text()]++
	}
	err = file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error closing file: %v", err)
		os.Exit(1)
	}
	for k, v := range words {
		fmt.Printf("%s\t%d\n", k, v)
	}
}