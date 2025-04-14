package greatestcommondivisorofstrings

import "testing"

func TestGcdOfStrings(t *testing.T) {
	tests := []struct {
		name     string
		str1     string
		str2     string
		expected string
	}{
		{
			name:     "example from README",
			str1:     "ABABAB",
			str2:     "ABAB",
			expected: "AB",
		},
		{
			name:     "no common divisor",
			str1:     "LEET",
			str2:     "CODE",
			expected: "",
		},
		{
			name:     "strings are equal",
			str1:     "ABCABC",
			str2:     "ABCABC",
			expected: "ABCABC",
		},
		{
			name:     "one string is multiple of another",
			str1:     "ABCABCABC",
			str2:     "ABC",
			expected: "ABC",
		},
		{
			name:     "different lengths with common divisor",
			str1:     "TAUXXTAUXXTAUXX",
			str2:     "TAUXXTAUXX",
			expected: "TAUXX",
		},
		{
			name:     "unicode characters",
			str1:     "приветпривет",
			str2:     "привет",
			expected: "привет",
		},
		{
			name:     "special characters",
			str1:     "!@#!@#!@#",
			str2:     "!@#!@#",
			expected: "!@#",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gcdOfStrings(tt.str1, tt.str2)
			if result != tt.expected {
				t.Errorf("gcdOfStrings(%q, %q) = %q, want %q", tt.str1, tt.str2, result, tt.expected)
			}
		})
	}
}
