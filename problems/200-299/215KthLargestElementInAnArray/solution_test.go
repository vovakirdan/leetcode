package kthlargestelementinanarray

import (
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
		{
			name:     "Single element",
			nums:     []int{10},
			k:        1,
			expected: 10,
		},
		{
			name:     "Two elements, first is k-th largest",
			nums:     []int{7, 8},
			k:        2,
			expected: 7,
		},
		{
			name:     "All elements equal",
			nums:     []int{5, 5, 5, 5, 5},
			k:        3,
			expected: 5,
		},
		{
			name:     "Descending order",
			nums:     []int{10, 9, 8, 7, 6},
			k:        5,
			expected: 6,
		},
		{
			name:     "Ascending order",
			nums:     []int{1, 2, 3, 4, 5},
			k:        2,
			expected: 4,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -2, -3, -4},
			k:        2,
			expected: -2,
		},
		{
			name:     "Mixed positive and negative numbers",
			nums:     []int{-10, 4, 2, 0, -3, 5},
			k:        3,
			expected: 2,
		},
		{
			name:     "Contains INT_MIN and INT_MAX",
			nums:     []int{-2147483648, 2147483647, 0},
			k:        2,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findKthLargest(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestFindKthLargestConstraints(t *testing.T) {
	t.Run("Minimum array size", func(t *testing.T) {
		nums := []int{42}
		k := 1
		result := findKthLargest(nums, k)
		if result != 42 {
			t.Errorf("Expected 42, got %d", result)
		}
	})

	t.Run("Maximum array size with same values", func(t *testing.T) {
		n := 100000 // 10^5
		nums := make([]int, n)
		for i := range nums {
			nums[i] = 10000
		}
		k := 1
		result := findKthLargest(nums, k)
		if result != 10000 {
			t.Errorf("Expected 10000, got %d", result)
		}
	})

	t.Run("Maximum array size with descending values", func(t *testing.T) {
		n := 100000
		nums := make([]int, n)
		for i := range nums {
			nums[i] = n - i
		}
		k := 1
		result := findKthLargest(nums, k)
		if result != 100000 {
			t.Errorf("Expected 100000, got %d", result)
		}
	})

	t.Run("Maximum array size with ascending values", func(t *testing.T) {
		n := 100000
		nums := make([]int, n)
		for i := range nums {
			nums[i] = i + 1
		}
		k := n
		result := findKthLargest(nums, k)
		if result != 1 {
			t.Errorf("Expected 1, got %d", result)
		}
	})
}
