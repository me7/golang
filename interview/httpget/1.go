// Process consumes URL strings from stdin.
// Process must GET each URL and count "Go" substring in response.
// After all URLs are processed, process should print total number of "Go"
// occurences, e.g.:
//
// $ echo -e 'https://golang.org\nhttps://golang.org\nhttps://golang.org' | go run 1.go
// Count for https://golang.org: 9
// Count for https://golang.org: 9
// Count for https://golang.org: 9
// Total: 27
//
// Each URL should be handled right after reading it, in parallel with
// reading other URLs. There should be no more that k=5 concurrent URL
// handlers. URL handlers must not spawn other goroutines, e.g. if k=1000
// and there are no URLs to handle, there should be no 1000 goroutines.
//
// Don't use global variables. Use only standard library.

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

const (
	MaxWorkers = 5
)

func get(url string) (count int) {

	// Set request timeout for one second.
	client := &http.Client{
		Timeout: 5e9,
	}

	// Issue GET.
	resp, err := client.Get(url)
	if err != nil {
		// Do error handling.
		return 0
	}
	defer resp.Body.Close()

	// Read response body into byte slice.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Do error handling.
		return 0
	}

	// Count substring number.
	return strings.Count(string(body), "Go")
}

// Convert stdin lines to a channel that emits "Go" count numbers.
func gen() <-chan int {

	// Set up wait group, semaphore and output channels.
	var wg sync.WaitGroup
	out := make(chan int)
	sem := make(chan int, MaxWorkers)

	go func() {
		// Read lines from stdin.
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			// Will block when sem is full - limit number of concurrent goroutines.
			sem <- 1
			wg.Add(1)
			go func(url string) {
				defer wg.Done()

				// Count "Go".
				count := get(url)
				fmt.Printf("Count for %s: %d\n", url, count)
				out <- count

				<-sem
			}(scanner.Text())
		}
		// Wait for all goroutines and close output channel.
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	fmt.Println("Reading from stdin...")

	sum := 0
	// Get the channel.
	in := gen()
	// Sum integers until the channel is closed.
	for n := range in {
		sum += n
	}

	// Print result.
	fmt.Println("Total:", sum)
}
