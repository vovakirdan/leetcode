package find_the_difference_of_two_arrays

import (
	"reflect"
	"testing"
)

func TestFindTheDifferenceOfTwoArrays(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected [][]int
	}{
		{
			name:     "Example from problem",
			nums1:    []int{1, 2, 3},
			nums2:    []int{2, 4, 6},
			expected: [][]int{{1, 3}, {4, 6}},
		},
		{
			name:     "Empty arrays",
			nums1:    []int{},
			nums2:    []int{},
			expected: [][]int{{}, {}},
		},
		{
			name:     "One empty array",
			nums1:    []int{1, 2, 3},
			nums2:    []int{},
			expected: [][]int{{1, 2, 3}, {}},
		},
		{
			name:     "Identical arrays",
			nums1:    []int{1, 2, 3},
			nums2:    []int{1, 2, 3},
			expected: [][]int{{}, {}},
		},
		{
			name:     "Maximum length arrays",
			nums1:    make([]int, 1000),
			nums2:    make([]int, 1000),
			expected: [][]int{{}, {}},
		},
		{
			name:     "Minimum and maximum values",
			nums1:    []int{-1000, 0, 1000},
			nums2:    []int{-1000, 1000},
			expected: [][]int{{0}, {}},
		},
		{
			name:     "Duplicate values in input",
			nums1:    []int{1, 1, 2, 2, 3, 3},
			nums2:    []int{2, 2, 4, 4},
			expected: [][]int{{1, 3}, {4}},
		},
		{
			name:     "All unique values",
			nums1:    []int{1, 3, 5, 7},
			nums2:    []int{2, 4, 6, 8},
			expected: [][]int{{1, 3, 5, 7}, {2, 4, 6, 8}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findDifference(tt.nums1, tt.nums2)

			// Sort the results for comparison since order doesn't matter
			sortSlice(result[0])
			sortSlice(result[1])
			sortSlice(tt.expected[0])
			sortSlice(tt.expected[1])

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("findDifference() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Helper function to sort slices for comparison
func sortSlice(s []int) {
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}
