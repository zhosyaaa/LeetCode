package main

import (
	"strconv"
)

func getDecimalValue(head *ListNode) int {
	var str string
	base := head

	for base != nil {
		str = str + strconv.Itoa(base.Val)
		base = base.Next
	}
	i, _ := strconv.ParseInt(str, 2, 64)
	return int(i)
}
