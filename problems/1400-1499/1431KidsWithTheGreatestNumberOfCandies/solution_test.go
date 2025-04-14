package kidswiththegreatestnumberofcandies

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		name         string
		candies      []int
		extraCandies int
		expected     []bool
	}{
		{
			name:         "example from README",
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			expected:     []bool{true, true, true, false, true},
		},
		{
			name:         "all kids can be greatest",
			candies:      []int{4, 2, 2, 4, 3},
			extraCandies: 2,
			expected:     []bool{true, true, true, true, true},
		},
		{
			name:         "no kids can be greatest except one",
			candies:      []int{1, 2, 3, 4, 6},
			extraCandies: 1,
			expected:     []bool{false, false, false, false, true},
		},
		{
			name:         "minimum array size",
			candies:      []int{1, 2},
			extraCandies: 1,
			expected:     []bool{true, true},
		},
		{
			name:         "maximum candy value",
			candies:      []int{100, 50, 75},
			extraCandies: 25,
			expected:     []bool{true, false, true},
		},
		{
			name:         "maximum extra candies",
			candies:      []int{1, 2, 3, 4, 5},
			extraCandies: 50,
			expected:     []bool{true, true, true, true, true},
		},
		{
			name:         "minimum candy value",
			candies:      []int{1, 1, 1, 1},
			extraCandies: 1,
			expected:     []bool{true, true, true, true},
		},
		{
			name:         "maximum array size",
			candies:      make([]int, 100), // All zeros
			extraCandies: 1,
			expected:     make([]bool, 100), // All false
		},
	}

	// Initialize the maximum array size test case
	for i := range tests[7].candies {
		tests[7].candies[i] = 1
		tests[7].expected[i] = true
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := kidsWithCandies(tt.candies, tt.extraCandies)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("kidsWithCandies(%v, %d) = %v, want %v", tt.candies, tt.extraCandies, result, tt.expected)
			}
		})
	}
}
