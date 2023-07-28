package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	lenth := 0
	current := head
	arr := make([]int, 0)
	for current != nil {
		lenth++
		arr = append(arr, current.Val)
		current = current.Next
	}
	arr = append(arr[:lenth-n], arr[lenth-n+1:]...)
	if len(arr) == 0 {
		return nil
	}
	head = &ListNode{Val: arr[0], Next: nil}
	res := head
	for i := 1; i < len(arr); i++ {
		temp := &ListNode{arr[i], nil}
		res.Next = temp
		res = res.Next
	}
	return head
}
