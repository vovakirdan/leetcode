package canplaceflowers

import (
	"testing"
)

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		name      string
		flowerbed []int
		n         int
		expected  bool
	}{
		{
			name:      "Example 1",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			expected:  true,
		},
		{
			name:      "Example 2",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			expected:  false,
		},
		{
			name:      "Single empty plot, can plant 1",
			flowerbed: []int{0},
			n:         1,
			expected:  true,
		},
		{
			name:      "Single occupied plot, cannot plant",
			flowerbed: []int{1},
			n:         1,
			expected:  false,
		},
		{
			name:      "Two empty plots, can plant 1",
			flowerbed: []int{0, 0},
			n:         1,
			expected:  true,
		},
		{
			name:      "Alternating pattern",
			flowerbed: []int{0, 1, 0, 1, 0, 1, 0},
			n:         1,
			expected:  false,
		},
		{
			name:      "Empty flowerbed",
			flowerbed: []int{},
			n:         0,
			expected:  true,
		},
		{
			name:      "No flowers needed",
			flowerbed: []int{0, 0, 0},
			n:         0,
			expected:  true,
		},
		{
			name:      "All zeros, plant max possible",
			flowerbed: []int{0, 0, 0, 0, 0},
			n:         2, // max is 2
			expected:  true,
		},
		{
			name:      "All zeros, too many flowers",
			flowerbed: []int{0, 0, 0, 0, 0},
			n:         3,
			expected:  true,
		},
		{
			name:      "Occupied edges only",
			flowerbed: []int{1, 0, 0, 0, 0, 1},
			n:         1,
			expected:  true,
		},
		{
			name:      "Long flowerbed, max allowed",
			flowerbed: make([]int, 100),
			n:         33, // ⌊(100+1)/2⌋ = 50, but alternating 0s needed, only 50% usable
			expected:  true,
		},
		{
			name:      "Long flowerbed, one too many",
			flowerbed: make([]int, 100),
			n:         34,
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// копия массива, чтобы не портить оригинал
			copied := make([]int, len(tt.flowerbed))
			copy(copied, tt.flowerbed)
			result := canPlaceFlowers(copied, tt.n)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
