package countingbits

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{
			name:     "example from README",
			n:        5,
			expected: []int{0, 1, 1, 2, 1, 2},
		},
		{
			name:     "n = 0",
			n:        0,
			expected: []int{0},
		},
		{
			name:     "n = 1",
			n:        1,
			expected: []int{0, 1},
		},
		{
			name:     "n = 2",
			n:        2,
			expected: []int{0, 1, 1},
		},
		{
			name:     "n = 3",
			n:        3,
			expected: []int{0, 1, 1, 2},
		},
		{
			name:     "n = 4",
			n:        4,
			expected: []int{0, 1, 1, 2, 1},
		},
		{
			name:     "n = 7 (all ones)",
			n:        7,
			expected: []int{0, 1, 1, 2, 1, 2, 2, 3},
		},
		{
			name:     "n = 8 (power of 2)",
			n:        8,
			expected: []int{0, 1, 1, 2, 1, 2, 2, 3, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countBits(tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("countBits(%d) = %v, want %v", tt.n, result, tt.expected)
			}
		})
	}
}
