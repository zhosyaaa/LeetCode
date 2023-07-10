package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	prev := &ListNode{-101, head}
	res := prev

	for prev.Next != nil {
		val := prev.Next.Val
		removed := false

		for prev.Next.Next != nil && prev.Next.Next.Val == val {
			prev.Next = prev.Next.Next
			removed = true
		}

		if removed {
			prev.Next = prev.Next.Next
			continue
		}
		prev = prev.Next
	}
	return res.Next
}
