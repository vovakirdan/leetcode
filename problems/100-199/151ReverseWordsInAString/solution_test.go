package reversewordsinastring

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Example 1 - simple",
			input:    "the sky is blue",
			expected: "blue is sky the",
		},
		{
			name:     "Example 2 - leading and trailing spaces",
			input:    "  hello world  ",
			expected: "world hello",
		},
		{
			name:     "Example 3 - multiple spaces between words",
			input:    "a good   example",
			expected: "example good a",
		},
		{
			name:     "Single word",
			input:    "word",
			expected: "word",
		},
		{
			name:     "Single word with spaces",
			input:    "   word   ",
			expected: "word",
		},
		{
			name:     "All spaces",
			input:    "     ",
			expected: "",
		},
		{
			name:     "Uppercase and lowercase mix",
			input:    "Go is FUN",
			expected: "FUN is Go",
		},
		{
			name:     "Numbers and words",
			input:    "123 456 abc",
			expected: "abc 456 123",
		},
		{
			name:     "Multiple spaces between and around words",
			input:    "   multiple   spaces  here   ",
			expected: "here spaces multiple",
		},
		{
			name:     "Only one character between spaces",
			input:    " a b c ",
			expected: "c b a",
		},
		{
			name:     "Long input",
			input:    "this is a very long input string with many words to test performance",
			expected: "performance test to words many with string input long very a is this",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseWords(tt.input)
			if result != tt.expected {
				t.Errorf("Input: %q — Expected: %q, got: %q", tt.input, tt.expected, result)
			}
		})
	}
}
