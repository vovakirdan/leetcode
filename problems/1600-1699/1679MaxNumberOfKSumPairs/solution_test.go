package maxnumberofksumpairs

import (
	"testing"
)

func TestMaxOperations(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1 - two perfect pairs",
			nums:     []int{1, 2, 3, 4},
			k:        5,
			expected: 2,
		},
		{
			name:     "Example 2 - only one pair possible",
			nums:     []int{3, 1, 3, 4, 3},
			k:        6,
			expected: 1,
		},
		{
			name:     "All pairs valid",
			nums:     []int{2, 2, 2, 2},
			k:        4,
			expected: 2,
		},
		{
			name:     "No valid pairs",
			nums:     []int{1, 2, 3, 4},
			k:        10,
			expected: 0,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			k:        5,
			expected: 0,
		},
		{
			name:     "One element only",
			nums:     []int{5},
			k:        5,
			expected: 0,
		},
		{
			name:     "Multiple duplicates forming multiple pairs",
			nums:     []int{1, 1, 2, 2, 3, 3, 4, 4},
			k:        5,
			expected: 4, // pairs: (1,4), (1,4), (2,3), (2,3)
		},
		{
			name:     "Pairs with repeated large numbers",
			nums:     []int{1000000000, 1000000000, 1000000000},
			k:        2000000000,
			expected: 1,
		},
		{
			name:     "Large input with no pairs",
			nums:     make([]int, 100000), // all 1s
			k:        3,
			expected: 0,
		},
		{
			name:     "Large input with many pairs",
			nums:     append(make([]int, 50000, 100000), make([]int, 50000)...), // 50000 1s + 50000 2s
			k:        3,
			expected: 0,
		},
		{
			name:     "Mixed pair overlap",
			nums:     []int{1, 5, 7, -1, 5},
			k:        6,
			expected: 2, // (1,5), (7,-1)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxOperations(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}
