package increasingtripletsubsequence

import "testing"

func TestIncreasingTriplet(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "empty array",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: false,
		},
		{
			name:     "two elements",
			nums:     []int{1, 2},
			expected: false,
		},
		{
			name:     "increasing triplet exists",
			nums:     []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "increasing triplet at the end",
			nums:     []int{5, 4, 3, 2, 1, 2, 3},
			expected: true,
		},
		{
			name:     "increasing triplet in the middle",
			nums:     []int{5, 1, 2, 3, 4},
			expected: true,
		},
		{
			name:     "no increasing triplet",
			nums:     []int{5, 4, 3, 2, 1},
			expected: false,
		},
		{
			name:     "equal elements",
			nums:     []int{1, 1, 1, 1, 1},
			expected: false,
		},
		{
			name:     "large numbers",
			nums:     []int{1000000, 2000000, 3000000},
			expected: true,
		},
		{
			name:     "negative numbers",
			nums:     []int{-5, -4, -3, -2, -1},
			expected: true,
		},
		{
			name:     "mixed positive and negative",
			nums:     []int{-1, 0, 1},
			expected: true,
		},
		{
			name:     "minimum possible triplet",
			nums:     []int{1, 2, 1, 3},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := increasingTriplet(tt.nums)
			if result != tt.expected {
				t.Errorf("increasingTriplet(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
