package uniquepaths

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name     string
		m        int
		n        int
		expected int
	}{
		{
			name:     "example from README",
			m:        3,
			n:        7,
			expected: 28,
		},
		{
			name:     "minimum grid size",
			m:        1,
			n:        1,
			expected: 1,
		},
		{
			name:     "single row",
			m:        1,
			n:        100,
			expected: 1,
		},
		{
			name:     "single column",
			m:        100,
			n:        1,
			expected: 1,
		},
		{
			name:     "square grid small",
			m:        2,
			n:        2,
			expected: 2,
		},
		{
			name:     "square grid medium",
			m:        3,
			n:        3,
			expected: 6,
		},
		{
			name:     "rectangle grid",
			m:        3,
			n:        2,
			expected: 3,
		},
		{
			name:     "large square grid",
			m:        10,
			n:        10,
			expected: 48620,
		},
		{
			name:     "large rectangle",
			m:        10,
			n:        20,
			expected: 6906900,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := uniquePaths(tt.m, tt.n)
			if result != tt.expected {
				t.Errorf("uniquePaths(%d, %d) = %d, want %d", tt.m, tt.n, result, tt.expected)
			}
		})
	}
}
