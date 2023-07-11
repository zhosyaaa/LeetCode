package main

func postorderTraversal(root *TreeNode) []int {
	output := make([]int, 0)
	if root == nil {
		return output
	}
	output = append(output, postorderTraversal(root.Left)...)
	output = append(output, postorderTraversal(root.Right)...)
	output = append(output, root.Val)
	return output
}
