package mergestringsalternately

import "testing"

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		name     string
		word1    string
		word2    string
		expected string
	}{
		{
			name:     "example from README",
			word1:    "ab",
			word2:    "pqrs",
			expected: "apbqrs",
		},
		{
			name:     "equal length strings",
			word1:    "abc",
			word2:    "def",
			expected: "adbecf",
		},
		{
			name:     "first string longer",
			word1:    "abcd",
			word2:    "ef",
			expected: "aebfcd",
		},
		{
			name:     "second string longer",
			word1:    "ab",
			word2:    "cdef",
			expected: "acbdef",
		},
		{
			name:     "empty first string",
			word1:    "",
			word2:    "xyz",
			expected: "xyz",
		},
		{
			name:     "empty second string",
			word1:    "abc",
			word2:    "",
			expected: "abc",
		},
		{
			name:     "both strings empty",
			word1:    "",
			word2:    "",
			expected: "",
		},
		{
			name:     "unicode characters",
			word1:    "привет",
			word2:    "мир",
			expected: "пмриирвет",
		},
		{
			name:     "special characters",
			word1:    "a!b@",
			word2:    "c#d$",
			expected: "ac!#bd@$",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeAlternately(tt.word1, tt.word2)
			if result != tt.expected {
				t.Errorf("mergeAlternately(%q, %q) = %q, want %q", tt.word1, tt.word2, result, tt.expected)
			}
		})
	}
}
