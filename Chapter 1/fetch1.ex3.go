// Fetch prints the content found at a URL.
// Don't forget to put http:// in the argument
package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url;
		}
		resp, err := http.Get(url)
		fmt.Printf("Status Code of the Answer: %s\n", resp.Status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
		bytes, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("Bytes copied: %d\n", bytes)
		fmt.Printf("Status Code of the Answer: %s\n", resp.Status)
	}
}