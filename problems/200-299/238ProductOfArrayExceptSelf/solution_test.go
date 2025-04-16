package productofarrayexceptself

import (
	"testing"
)

func TestProductOfArrayExceptSelf(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "Example 2",
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
		{
			name:     "Minimum length array",
			nums:     []int{2, 3},
			expected: []int{3, 2},
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "Single zero",
			nums:     []int{1, 2, 0, 4},
			expected: []int{0, 0, 8, 0},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-2, -3, -4},
			expected: []int{12, 8, 6},
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{-2, 3, -4, 5},
			expected: []int{-60, 40, -30, 24},
		},
		{
			name:     "Maximum value",
			nums:     []int{30, 30, 30},
			expected: []int{900, 900, 900},
		},
		{
			name:     "Minimum value",
			nums:     []int{-30, -30, -30},
			expected: []int{900, 900, 900},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := productExceptSelf(tt.nums)
			if len(result) != len(tt.expected) {
				t.Errorf("Expected length %d, got %d", len(tt.expected), len(result))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("At index %d: expected %d, got %d", i, tt.expected[i], result[i])
				}
			}
		})
	}
}

func TestProductOfArrayExceptSelfConstraints(t *testing.T) {
	// Test minimum length
	t.Run("Minimum length", func(t *testing.T) {
		nums := []int{1, 2}
		result := productExceptSelf(nums)
		if len(result) != 2 {
			t.Errorf("Expected length 2, got %d", len(result))
		}
	})

	// Test maximum value constraint
	t.Run("Maximum value", func(t *testing.T) {
		nums := make([]int, 100000) // 10^5 elements
		for i := range nums {
			nums[i] = 30 // Maximum allowed value
		}
		result := productExceptSelf(nums)
		if len(result) != len(nums) {
			t.Errorf("Expected length %d, got %d", len(nums), len(result))
		}
	})

	// Test minimum value constraint
	t.Run("Minimum value", func(t *testing.T) {
		nums := make([]int, 100000) // 10^5 elements
		for i := range nums {
			nums[i] = -30 // Minimum allowed value
		}
		result := productExceptSelf(nums)
		if len(result) != len(nums) {
			t.Errorf("Expected length %d, got %d", len(nums), len(result))
		}
	})
}
