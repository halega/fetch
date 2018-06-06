package main

import (
	"fmt"
	"log"
	"os"

	"github.com/valyala/fasthttp"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: fetch <url>\n")
		os.Exit(1)
	}

	c := &fasthttp.Client{}
	statusCode, body, err := c.Get(nil, os.Args[1])
	if err != nil {
		log.Fatalf("Error when loading localhost:8080: %s", err)
	}
	if statusCode != fasthttp.StatusOK {
		log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	}
	fmt.Printf("Body's length: %d\n", len(body))
	fmt.Println(string(body))
}
