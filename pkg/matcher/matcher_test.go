package matcher

import (
	"testing"
)

func TestMatcherStruct_Match(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		data     []string
		expected []string
	}{
		{
			name:     "Empty pattern and data",
			pattern:  "",
			data:     []string{},
			expected: []string{},
		},
		{
			name:     "Pattern matches some data",
			pattern:  "foo",
			data:     []string{"foo", "bar", "baz"},
			expected: []string{"foo"},
		},
		{
			name:     "Pattern matches all data",
			pattern:  "abc",
			data:     []string{"abc", "abc", "abc"},
			expected: []string{"abc", "abc", "abc"},
		},
		{
			name:     "Pattern does not match any data",
			pattern:  "xyz",
			data:     []string{"abc", "def", "ghi"},
			expected: []string{},
		},
	}

	matcher := NewMatcherStruct()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := matcher.Match(test.pattern, test.data)

			if len(result) != len(test.expected) {
				t.Errorf("Expected %d matches, but got %d", len(test.expected), len(result))
			}

			for i, match := range result {
				if match != test.expected[i] {
					t.Errorf("Expected match %s, but got %s", test.expected[i], match)
				}
			}
		})
	}
}

func BenchmarkMatcherStruct_Match(b *testing.B) {
	pattern := "foo"
	data := []string{"foo", "bar", "baz"}
	matcher := NewMatcherStruct()

	for i := 0; i < b.N; i++ {
		matcher.Match(pattern, data)
	}
}