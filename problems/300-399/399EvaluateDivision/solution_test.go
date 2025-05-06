package evaluatedivision

import (
	"math"
	"testing"
)

func floatEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-5
}

func sliceFloatEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !floatEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestCalcEquation(t *testing.T) {
	tests := []struct {
		name      string
		equations [][]string
		values    []float64
		queries   [][]string
		expected  []float64
	}{
		{
			name:      "Example 1",
			equations: [][]string{{"a", "b"}, {"b", "c"}},
			values:    []float64{2.0, 3.0},
			queries: [][]string{
				{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"},
			},
			expected: []float64{6.0, 0.5, -1.0, 1.0, -1.0},
		},
		{
			name:      "Example 2",
			equations: [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
			values:    []float64{1.5, 2.5, 5.0},
			queries: [][]string{
				{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"},
			},
			expected: []float64{3.75, 0.4, 5.0, 0.2},
		},
		{
			name:      "Example 3",
			equations: [][]string{{"a", "b"}},
			values:    []float64{0.5},
			queries: [][]string{
				{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"},
			},
			expected: []float64{0.5, 2.0, -1.0, -1.0},
		},
		{
			name:      "Self division",
			equations: [][]string{{"a", "b"}, {"b", "c"}},
			values:    []float64{2.0, 3.0},
			queries:   [][]string{{"c", "c"}, {"b", "b"}, {"a", "a"}},
			expected:  []float64{1.0, 1.0, 1.0},
		},
		{
			name:      "Disconnected components",
			equations: [][]string{{"a", "b"}, {"c", "d"}},
			values:    []float64{2.0, 3.0},
			queries: [][]string{
				{"a", "d"}, {"b", "c"},
			},
			expected: []float64{-1.0, -1.0},
		},
		{
			name:      "Cycle check",
			equations: [][]string{{"a", "b"}, {"b", "c"}, {"c", "a"}},
			values:    []float64{2.0, 3.0, 0.5},
			queries:   [][]string{{"a", "c"}, {"c", "b"}, {"b", "a"}},
			expected:  []float64{2, float64(1) / 3, 0.5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calcEquation(tt.equations, tt.values, tt.queries)
			if !sliceFloatEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
