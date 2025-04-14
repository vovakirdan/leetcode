package uniquenumberofoccurrences

import "testing"

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected bool
	}{
		{
			name:     "example from README",
			arr:      []int{1, 2, 2, 1, 1, 3},
			expected: true,
		},
		{
			name:     "empty array",
			arr:      []int{},
			expected: true,
		},
		{
			name:     "single element",
			arr:      []int{1},
			expected: true,
		},
		{
			name:     "all same elements",
			arr:      []int{1, 1, 1, 1},
			expected: true,
		},
		{
			name:     "duplicate frequencies",
			arr:      []int{1, 1, 2, 2, 3, 3},
			expected: false,
		},
		{
			name:     "multiple unique frequencies",
			arr:      []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			expected: true,
		},
		{
			name:     "negative numbers with unique frequencies",
			arr:      []int{-1, -1, -2, -2, -2, -3, -3, -3, -3},
			expected: true,
		},
		{
			name:     "mixed positive and negative with duplicate frequencies",
			arr:      []int{1, -1, 2, -2, 3, -3},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := uniqueOccurrences(tt.arr)
			if result != tt.expected {
				t.Errorf("uniqueOccurrences(%v) = %v, want %v", tt.arr, result, tt.expected)
			}
		})
	}
}
