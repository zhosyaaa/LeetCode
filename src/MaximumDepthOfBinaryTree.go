package main

//func main() {
//	p := &TreeNode{Val: 1}
//	p.Left = &TreeNode{Val: 2}
//	fmt.Println(maxDepth(p))
//}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Right), maxDepth(root.Left))
}

//
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
