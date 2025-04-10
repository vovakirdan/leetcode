package binarytreerightsideview

import (
	"testing"
)

func TestRightSideView(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name: "single node",
			root: &TreeNode{
				Val: 1,
			},
			expected: []int{1},
		},
		{
			name: "complete binary tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: []int{1, 3, 7},
		},
		{
			name: "left skewed tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "right skewed tree",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "unbalanced tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			expected: []int{1, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RightSideView(tt.root)
			if len(result) != len(tt.expected) {
				t.Errorf("expected length %d, got %d", len(tt.expected), len(result))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("at index %d: expected %d, got %d", i, tt.expected[i], result[i])
				}
			}
		})
	}
}
