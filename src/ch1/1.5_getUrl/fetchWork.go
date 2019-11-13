package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		var prefix_url string
		if prefix := strings.HasPrefix(url, "http://"); prefix != true {
			prefix_url = "http://" + url
		}

		fmt.Println("url: ", prefix_url)

		resp, err := http.Get(prefix_url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}

		fmt.Println("resp.Status: ", resp.Status)

		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			log.Fatal(err)
		}
	}
}
