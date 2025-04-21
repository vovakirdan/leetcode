package binarytreezigzaglevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		size := len(queue)
		row := make([]int, size)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if level%2 == 0 {
				row[i] = node.Val
			} else {
				row[size-1-i] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, row)
		level++
	}

	return res
}
