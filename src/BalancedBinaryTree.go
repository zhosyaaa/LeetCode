package main

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	dif := helps(root.Right) - helps(root.Left)
	if dif <= 1 && dif >= -1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}

func helps(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(helps(root.Right), helps(root.Left))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
