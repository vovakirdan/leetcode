package findpeakelement

import (
	"slices"
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		name        string
		nums        []int
		expected_in []int
	}{
		{
			name:        "single peak in middle",
			nums:        []int{1, 2, 3, 1},
			expected_in: []int{2},
		},
		{
			name:        "peak at start",
			nums:        []int{3, 2, 1},
			expected_in: []int{0},
		},
		{
			name:        "peak at end",
			nums:        []int{1, 2, 3},
			expected_in: []int{2},
		},
		{
			name:        "multiple peaks",
			nums:        []int{1, 2, 1, 3, 5, 6, 4},
			expected_in: []int{5, 1}, // 5 and 1 are both peaks
		},
		{
			name:        "single element",
			nums:        []int{1},
			expected_in: []int{0},
		},
		{
			name:        "two elements ascending",
			nums:        []int{1, 2},
			expected_in: []int{1},
		},
		{
			name:        "two elements descending",
			nums:        []int{2, 1},
			expected_in: []int{0},
		},
		{
			name:        "peak with distinct neighbors",
			nums:        []int{1, 2, 1, 3, 4},
			expected_in: []int{1, 3, 4}, // 1 and 3 are both peaks
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findPeakElement(tt.nums)
			if !slices.Contains(tt.expected_in, result) {
				t.Errorf("findPeakElement(%v) = %v, want one of %v", tt.nums, result, tt.expected_in)
			}
		})
	}
}
