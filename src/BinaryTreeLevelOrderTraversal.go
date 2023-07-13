package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	level := 0
	current := make([]*TreeNode, 1)
	current[0] = root
	for len(current) > 0 {
		curLen := len(current)
		for i := 0; i < curLen; i++ {
			node := current[0]
			current = current[1:]

			if len(res) <= level {
				res = append(res, []int{node.Val})
			} else {
				res[level] = append(res[level], node.Val)
			}

			if node.Left != nil {
				current = append(current, node.Left)
			}
			if node.Right != nil {
				current = append(current, node.Right)
			}
		}
		level++
	}
	return res
}
