package matcher

import (
	"strings"
	"sync"

	"github.com/fatih/color"
)

// Create a mutex to synchronize access to the counter variable
var mu sync.Mutex

// Create a WaitGroup to wait for all the goroutines to finish
var wg sync.WaitGroup

// Create an interface for the matcher with a single method, Match
type Matcher interface {
	Match(string, []string) ([]string, int)
}

// Create a struct to hold the matcher
type MatcherStruct struct {
	Matcher
}

// Create a function to return a new MatcherStruct
func NewMatcherStruct() *MatcherStruct {
	return &MatcherStruct{}
}

// Create the Match method for the MatcherStruct
func (m *MatcherStruct) Match(pattern string, data []string) ([]string, int) {
	resultChan := make(chan string)
	var counter int

	for _, line := range data {
		wg.Add(1)
		go func(line string) {
			if strings.Contains(line, pattern) {
				mu.Lock()
				counter += 1
				mu.Unlock()
				// Send the data the channel with the pattern highlighted
				cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
				match := strings.ReplaceAll(string(line), pattern, cyan(pattern))
				resultChan <- match
			}
			wg.Done()
		}(line)
	}

	// Wait for all the goroutines to finish and close the channel
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect the results from the channel and return them
	results := make([]string, 0)
	for result := range resultChan {
		results = append(results, result)
	}

	return results, counter
}
