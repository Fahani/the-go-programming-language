// Dup2 prints the count and text of lines that appear more than once in the input. It reads from stdin or from alist of named files
// Mofify dump2 to print the names of all files in which each duplicated line occurs.
package main

import(
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		for fileName, repetitions := range n {
			fmt.Printf("%s\t%d\t%s\n",fileName, repetitions, line)
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			temp := make(map[string]int)
			temp[fileName]++
			counts[input.Text()]=temp
		} else {
			counts[input.Text()][fileName]++
		}
	}
}