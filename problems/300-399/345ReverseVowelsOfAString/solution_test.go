package reversevowelsofastring

import (
	"testing"
)

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Example 1 - Mixed case",
			input:    "IceCreAm",
			expected: "AceCreIm",
		},
		{
			name:     "Example 2 - Lowercase",
			input:    "leetcode",
			expected: "leotcede",
		},
		{
			name:     "Only vowels lowercase",
			input:    "aeiou",
			expected: "uoiea",
		},
		{
			name:     "Only vowels uppercase",
			input:    "AEIOU",
			expected: "UOIEA",
		},
		{
			name:     "Mixed vowels",
			input:    "aEiOu",
			expected: "uOiEa",
		},
		{
			name:     "No vowels",
			input:    "bcdfg",
			expected: "bcdfg",
		},
		{
			name:     "One vowel only",
			input:    "a",
			expected: "a",
		},
		{
			name:     "Vowel at start and end",
			input:    "aBCDe",
			expected: "eBCDa",
		},
		{
			name:     "Palindrome with vowels",
			input:    "madam",
			expected: "madam",
		},
		{
			name:     "Alternating vowels and consonants",
			input:    "a-b-c-d-e",
			expected: "e-b-c-d-a",
		},
		{
			name:     "Long string no vowels",
			input:    "bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ",
			expected: "bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ",
		},
		{
			name:     "Long string with repeated vowels",
			input:    "aaaeeeiiiooouuu",
			expected: "uuuoooiiieeeaaa",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "String with non-letters",
			input:    "1a!e?o",
			expected: "1o!e?a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseVowels(tt.input)
			if result != tt.expected {
				t.Errorf("Input: %q — Expected: %q, got: %q", tt.input, tt.expected, result)
			}
		})
	}
}
