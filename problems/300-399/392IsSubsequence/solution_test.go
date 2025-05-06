package issubsequence

import (
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "Example 1 - s is subsequence",
			s:        "abc",
			t:        "ahbgdc",
			expected: true,
		},
		{
			name:     "Example 2 - s is not subsequence",
			s:        "axc",
			t:        "ahbgdc",
			expected: false,
		},
		{
			name:     "Empty s is always subsequence",
			s:        "",
			t:        "anystring",
			expected: true,
		},
		{
			name:     "Empty t and non-empty s",
			s:        "a",
			t:        "",
			expected: false,
		},
		{
			name:     "Both s and t are empty",
			s:        "",
			t:        "",
			expected: true,
		},
		{
			name:     "s equals t",
			s:        "abc",
			t:        "abc",
			expected: true,
		},
		{
			name:     "s is longer than t",
			s:        "abcdef",
			t:        "abc",
			expected: false,
		},
		{
			name:     "s is subsequence with skips",
			s:        "ace",
			t:        "abcde",
			expected: true,
		},
		{
			name:     "s with repeated letters",
			s:        "aaa",
			t:        "aaabaa",
			expected: true,
		},
		{
			name:     "t with repeated but out of order",
			s:        "abc",
			t:        "acb",
			expected: false,
		},
		{
			name:     "Long t with scattered s",
			s:        "xyz",
			t:        "aaxxxbyyyczzz",
			expected: true,
		},
		{
			name:     "s present but order broken",
			s:        "abc",
			t:        "bac",
			expected: false,
		},
		{
			name:     "Long t no match",
			s:        "abc",
			t:        string(make([]byte, 10000)), // all zeros
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSubsequence(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("s: %q, t: %q — expected %v, got %v", tt.s, tt.t, tt.expected, result)
			}
		})
	}
}
