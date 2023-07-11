package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left != nil && root.Right != nil {
		return 1 + min(minDepth(root.Left), minDepth(root.Right))
	}
	if root.Left != nil {
		return 1 + minDepth(root.Left)
	} else {
		return 1 + minDepth(root.Right)
	}
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
