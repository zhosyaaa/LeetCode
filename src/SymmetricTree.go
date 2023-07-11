package main

//func main() {
//	p := &TreeNode{Val: 1}
//	p.Left = &TreeNode{Val: 2}
//	q := &TreeNode{Val: 1}
//	q.Right = &TreeNode{Val: 2}
//	fmt.Println(isSameTree(p, q))
//}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

func helper(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}
