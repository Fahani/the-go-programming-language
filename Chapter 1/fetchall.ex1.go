// Fetchall fetches URLs in parallel and reports their times and sizes
package main

import (
	"fmt"
	"os"
	"time"
	"net/http"
	"io"
	"strings"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // Starts a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan <- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	f, err := os.Create(strings.Replace(strings.Replace(url,"/", ".",-1),":", ".",-1)+".txt")
	if err != nil {
		f.Close()
		resp.Body.Close()
		ch <- fmt.Sprintf("while creating file %s: %v", url, err)
		return
	}
	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close()
	f.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}