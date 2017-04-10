package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

/**
Exercise 1.9
*/
func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Print(resp.Status)
	}

}
