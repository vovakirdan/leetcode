package binarytreezigzaglevelordertraversal

import (
	"reflect"
	"testing"
)

func TestBinaryTreeZigzagLevelOrderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: [][]int{},
		},
		{
			name: "Single node",
			root: &TreeNode{
				Val: 1,
			},
			expected: [][]int{{1}},
		},
		{
			name: "Left skewed tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name: "Right skewed tree",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name: "Complete binary tree",
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
			expected: [][]int{{1}, {3, 2}, {4, 5, 6, 7}},
		},
		{
			name: "Unbalanced tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: [][]int{{1}, {3, 2}, {4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := zigzagLevelOrder(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", result, tt.expected)
			}
		})
	}
}
