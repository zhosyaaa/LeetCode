package main

//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func preorderTraversal(root *TreeNode) []int {
	output := make([]int, 0)
	if root == nil {
		return output
	}
	output = append(output, root.Val)
	output = append(output, preorderTraversal(root.Left)...)
	output = append(output, preorderTraversal(root.Right)...)
	return output
}
