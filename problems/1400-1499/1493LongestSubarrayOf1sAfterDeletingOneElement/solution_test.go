package longestsubarrayof1safterdeletingoneelement

import "testing"

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "all ones",
			nums:     []int{1, 1, 1, 1, 1},
			expected: 4, // must delete one element
		},
		{
			name:     "all zeros",
			nums:     []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "single element array",
			nums:     []int{1},
			expected: 0, // must delete one element
		},
		{
			name:     "example from problem",
			nums:     []int{1, 1, 0, 1},
			expected: 3,
		},
		{
			name:     "multiple zeros",
			nums:     []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			expected: 5,
		},
		{
			name:     "zeros at start and end",
			nums:     []int{0, 1, 1, 1, 1, 0},
			expected: 4,
		},
		{
			name:     "alternating ones and zeros",
			nums:     []int{1, 0, 1, 0, 1, 0, 1},
			expected: 2,
		},
		{
			name:     "large array",
			nums:     make([]int, 100000), // all zeros
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestSubarray(tt.nums)
			if result != tt.expected {
				t.Errorf("longestSubarray() = %v, want %v", result, tt.expected)
			}
		})
	}
}
