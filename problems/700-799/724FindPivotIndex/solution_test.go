package findpivotindex

import "testing"

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "single element",
			nums:     []int{1},
			expected: 0,
		},
		{
			name:     "pivot at the beginning",
			nums:     []int{2, 1, -1},
			expected: 0,
		},
		{
			name:     "pivot in the middle",
			nums:     []int{1, 7, 3, 6, 5, 6},
			expected: 3,
		},
		{
			name:     "pivot at the end",
			nums:     []int{-1, -1, -1, 0, 1, 1},
			expected: 0,
		},
		{
			name:     "no pivot exists",
			nums:     []int{1, 2, 3},
			expected: -1,
		},
		{
			name:     "all zeros",
			nums:     []int{0, 0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -1, -1, -1, -1},
			expected: 2,
		},
		{
			name:     "mixed positive and negative",
			nums:     []int{-1, 1, -1, 1, -1, 1},
			expected: -1,
		},
		{
			name:     "maximum length array",
			nums:     make([]int, 10000),
			expected: 0,
		},
		{
			name:     "minimum value elements",
			nums:     []int{-1000, -1000, -1000},
			expected: 1,
		},
		{
			name:     "maximum value elements",
			nums:     []int{1000, 1000, 1000},
			expected: 1,
		},
		{
			name:     "alternating large numbers",
			nums:     []int{1000, -1000, 1000, -1000},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pivotIndex(tt.nums)
			if result != tt.expected {
				t.Errorf("pivotIndex(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
