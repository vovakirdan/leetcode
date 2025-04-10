package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "basic case",
			nums:     []int{2, 2, 1},
			expected: 1,
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "multiple pairs",
			nums:     []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -1, -2},
			expected: -2,
		},
		{
			name:     "mixed positive and negative",
			nums:     []int{-1, 1, -1, 1, 0},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := singleNumber(tt.nums)
			if result != tt.expected {
				t.Errorf("singleNumber(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
