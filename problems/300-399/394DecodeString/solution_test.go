package decodestring

import (
	"strings"
	"testing"
)

func TestDecodeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Example 1",
			input:    "3[a]2[bc]",
			expected: "aaabcbc",
		},
		{
			name:     "Example 2",
			input:    "3[a2[c]]",
			expected: "accaccacc",
		},
		{
			name:     "Example 3",
			input:    "2[abc]3[cd]ef",
			expected: "abcabccdcdcdef",
		},
		{
			name:     "Single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "Single repetition",
			input:    "2[a]",
			expected: "aa",
		},
		{
			name:     "Nested brackets",
			input:    "3[a2[b3[c]]]",
			expected: "abcccbcccabcccbcccabcccbccc",
		},
		{
			name:     "Maximum repetition",
			input:    "300[a]",
			expected: strings.Repeat("a", 300),
		},
		{
			name:     "Multiple nested levels",
			input:    "2[3[a]2[b]]",
			expected: "aaabbaaabb",
		},
		{
			name:     "Complex nested structure",
			input:    "2[2[ab]3[cd]]",
			expected: "ababcdcdcdababcdcdcd",
		},
		{
			name:     "Empty string inside brackets",
			input:    "2[]",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := decodeString(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}
