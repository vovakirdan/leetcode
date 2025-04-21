package validpalindromeii

import "testing"

func TestValidPalindromeII(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "Example 1 - aba",
			s:        "aba",
			expected: true,
		},
		{
			name:     "Example 2 - abca",
			s:        "abca",
			expected: true,
		},
		{
			name:     "Example 3 - abc",
			s:        "abc",
			expected: false,
		},
		{
			name:     "Empty string",
			s:        "",
			expected: true,
		},
		{
			name:     "Single character",
			s:        "a",
			expected: true,
		},
		{
			name:     "Two different characters",
			s:        "ab",
			expected: true,
		},
		{
			name:     "Long palindrome with one mismatch",
			s:        "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga",
			expected: true,
		},
		{
			name:     "String requiring two deletions",
			s:        "abcde",
			expected: false,
		},
		{
			name:     "String with multiple mismatches",
			s:        "abccbda",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validPalindrome(tt.s)
			if result != tt.expected {
				t.Errorf("validPalindrome(%q) = %v, want %v", tt.s, result, tt.expected)
			}
		})
	}
}
