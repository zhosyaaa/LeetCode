package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	output := make([]int, 0)
	if root == nil {
		return output
	}
	output = append(output, inorderTraversal(root.Left)...)
	output = append(output, root.Val)
	output = append(output, inorderTraversal(root.Right)...)
	return output
}
