package binarytreerightsideview

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func RightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	result := []int{}

	for len(queue) > 0 {
		level_size := len(queue)

		for i := 0; i < level_size; i++ {
			node := queue[0]
			queue = queue[1:]

			if i == level_size-1 {
				result = append(result, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}
