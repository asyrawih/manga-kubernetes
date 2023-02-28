package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/iain17/go-cfscrape"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s url\n", os.Args[0])
	}

	url := os.Args[1]

	// First get will incur cloudflare challenge
	resp, err := cfscrape.Get(url)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()

	// Try getting the same URL again; none of these goroutines should incur a challenge
	var wg sync.WaitGroup

	const numGoroutines = 5
	results := make(chan []byte, numGoroutines)
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			// First get will incur cloudflare challenge
			resp, err := cfscrape.Get(url)
			if err != nil {
				panic(err)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			resp.Body.Close()

			results <- body
			wg.Done()
		}()
	}

	wg.Wait()

	for i := 0; i < numGoroutines; i++ {
		fmt.Println(string(<-results))
	}

}
