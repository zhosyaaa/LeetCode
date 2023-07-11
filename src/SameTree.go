package main

//	func main() {
//		p := &TreeNode{Val: 1}
//		p.Left = &TreeNode{Val: 2}
//		q := &TreeNode{Val: 1}
//		q.Right = &TreeNode{Val: 2}
//		fmt.Println(isSameTree(p, q))
//	}
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
