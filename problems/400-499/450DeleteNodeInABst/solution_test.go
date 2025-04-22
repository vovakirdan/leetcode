package deletenodeinabst

import (
	"reflect"
	"testing"
)

// Вспомогательная функция для построения BST из массива (insert)
func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}
	return root
}

// Вспомогательная функция для преобразования дерева в срез in-order
func inOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, res)
	*res = append(*res, root.Val)
	inOrder(root.Right, res)
}

// Тесты
func TestDeleteNode(t *testing.T) {
	tests := []struct {
		name         string
		initial      []int
		deleteKey    int
		expectedInOrder []int
	}{
		{
			name:         "Delete leaf node",
			initial:      []int{5, 3, 6, 2, 4},
			deleteKey:    2,
			expectedInOrder: []int{3, 4, 5, 6}, // in-order: [3, 4, 5, 6]
		},
		{
			name:         "Delete node with one child",
			initial:      []int{5, 3, 6, 2},
			deleteKey:    3,
			expectedInOrder: []int{2, 5, 6},
		},
		{
			name:         "Delete node with two children",
			initial:      []int{5, 3, 6, 2, 4},
			deleteKey:    3,
			expectedInOrder: []int{2, 4, 5, 6},
		},
		{
			name:         "Delete root node",
			initial:      []int{5, 3, 6, 2, 4},
			deleteKey:    5,
			expectedInOrder: []int{2, 3, 4, 6},
		},
		{
			name:         "Delete non-existent key",
			initial:      []int{5, 3, 6},
			deleteKey:    100,
			expectedInOrder: []int{3, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var root *TreeNode
			for _, val := range tt.initial {
				root = insert(root, val)
			}
			root = deleteNode(root, tt.deleteKey)

			var result []int
			inOrder(root, &result)

			if !reflect.DeepEqual(result, tt.expectedInOrder) {
				t.Errorf("Expected in-order %v, got %v", tt.expectedInOrder, result)
			}
		})
	}
}
