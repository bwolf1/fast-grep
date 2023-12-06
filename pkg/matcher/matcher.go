package matcher

import (
	"strings"

	"github.com/fatih/color"
)

// Create an interface for the matcher with a single method, Match
type Matcher interface {
	Match([]byte) bool
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
func (m *MatcherStruct) Match(pattern string, data string) string {
	if !strings.Contains(data, pattern) {
		// Mimic the behavior of GNU grep when no match is found
		return ""
	}

	// Return the data with the pattern highlighted
	cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
	match := strings.ReplaceAll(string(data), pattern, cyan(pattern))

	return match
}
