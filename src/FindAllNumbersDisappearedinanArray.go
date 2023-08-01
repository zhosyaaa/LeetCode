package main

func findDisappearedNumbers(nums []int) []int {
	temp := make([]int, 0)
	maps := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		maps[nums[i]] = 1
	}
	for i := 1; i <= len(nums); i++ {
		if maps[i] == 0 {
			temp = append(temp, i)
		}
	}
	return temp
}

// Remove Linked List Elements
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	ans, temp := head, head
	for temp != nil {
		if temp.Val == val {
			ans.Next = temp.Next
		} else {
			ans = temp
		}
		temp = temp.Next
	}
	return head
}
