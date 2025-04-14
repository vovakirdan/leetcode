package removingstarsfromastring

import "testing"

func TestRemoveStars(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic test",
			input:    "leet**cod*e",
			expected: "lecoe",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "no stars",
			input:    "leetcode",
			expected: "leetcode",
		},
		{
			name:     "all stars",
			input:    "****",
			expected: "",
		},
		{
			name:     "stars at the end",
			input:    "leetcode***",
			expected: "leetc",
		},
		{
			name:     "stars at the beginning",
			input:    "***leetcode",
			expected: "leetcode",
		},
		{
			name:     "alternating stars",
			input:    "l*e*e*t*c*o*d*e",
			expected: "e",
		},
		{
			name:     "multiple stars in sequence",
			input:    "leet****code",
			expected: "code",
		},
		{
			name:     "unicode characters",
			input:    "привет*мир*",
			expected: "привеми",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeStars(tt.input)
			if result != tt.expected {
				t.Errorf("removeStars(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
