package maximumsubsequencescore

import (
	"testing"
)

func TestMaxScore(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		k        int
		expected int64
	}{
		{
			name:     "Example 1",
			nums1:    []int{1, 3, 3, 2},
			nums2:    []int{2, 1, 3, 4},
			k:        3,
			expected: 12,
		},
		{
			name:     "Example 2",
			nums1:    []int{4, 2, 3, 1, 1},
			nums2:    []int{7, 5, 10, 9, 6},
			k:        1,
			expected: 30,
		},
		{
			name:     "All nums1 are zeros",
			nums1:    []int{0, 0, 0, 0},
			nums2:    []int{1, 2, 3, 4},
			k:        2,
			expected: 0,
		},
		{
			name:     "All nums2 are same",
			nums1:    []int{1, 2, 3, 4},
			nums2:    []int{5, 5, 5, 5},
			k:        3,
			expected: int64(9 * 5),
		},
		{
			name:     "Minimum k = 1, choose max nums1 * corresponding nums2",
			nums1:    []int{1, 10, 5},
			nums2:    []int{2, 2, 10},
			k:        1,
			expected: 50, // choose index 2
		},
		{
			name:     "Large k = len(nums1)",
			nums1:    []int{5, 1, 3},
			nums2:    []int{4, 4, 4},
			k:        3,
			expected: int64((5 + 1 + 3) * 4),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxScore(tt.nums1, tt.nums2, tt.k)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func BenchmarkMaxScore(b *testing.B) {
	n := 100_000
	nums1 := make([]int, n)
	nums2 := make([]int, n)
	for i := 0; i < n; i++ {
		nums1[i] = i % 1000
		nums2[i] = 100_000 - i
	}
	k := 50_000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = maxScore(nums1, nums2, k)
	}
}
