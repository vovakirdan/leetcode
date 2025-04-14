package rottingoranges

import "testing"

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "example from README",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			expected: 4,
		},
		{
			name: "all oranges already rotten",
			grid: [][]int{
				{2, 2, 2},
				{2, 2, 2},
				{2, 2, 2},
			},
			expected: 0,
		},
		{
			name: "impossible to rot all oranges",
			grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			expected: -1,
		},
		{
			name: "single fresh orange",
			grid: [][]int{
				{2, 1},
			},
			expected: 1,
		},
		{
			name: "single rotten orange",
			grid: [][]int{
				{2},
			},
			expected: 0,
		},
		{
			name: "empty grid",
			grid: [][]int{
				{},
			},
			expected: 0,
		},
		{
			name: "large grid with multiple sources",
			grid: [][]int{
				{2, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 2},
			},
			expected: 3,
		},
		{
			name: "grid with obstacles",
			grid: [][]int{
				{2, 1, 1, 0},
				{1, 1, 0, 1},
				{0, 1, 1, 1},
				{1, 0, 1, 2},
			},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the grid for each test to avoid modifying the original
			gridCopy := make([][]int, len(tt.grid))
			for i := range tt.grid {
				gridCopy[i] = make([]int, len(tt.grid[i]))
				copy(gridCopy[i], tt.grid[i])
			}

			result := orangesRotting(gridCopy)
			if result != tt.expected {
				t.Errorf("orangesRotting(%v) = %d, want %d", tt.grid, result, tt.expected)
			}
		})
	}
}
