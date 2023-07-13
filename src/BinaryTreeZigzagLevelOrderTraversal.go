package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	level := 0
	current := make([]*TreeNode, 1)
	current[0] = root
	isLeft := false
	for len(current) > 0 {
		curLen := len(current)
		levelRes := []int{}
		for i := 0; i < curLen; i++ {
			node := current[0]
			current = current[1:]

			levelRes = append(levelRes, node.Val)

			if node.Left != nil {
				current = append(current, node.Left)
			}
			if node.Right != nil {
				current = append(current, node.Right)
			}
		}

		// Reverse the level result if necessary
		if isLeft {
			reverseSlice(levelRes)
		}

		res = append(res, levelRes)
		isLeft = !isLeft
		level++
	}

	return res
}

func reverseSlice(slice []int) {
	left := 0
	right := len(slice) - 1
	for left < right {
		slice[left], slice[right] = slice[right], slice[left]
		left++
		right--
	}
}
