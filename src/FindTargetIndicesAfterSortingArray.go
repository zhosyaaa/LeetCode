package main

import (
	"sort"
)

func targetIndices(nums []int, target int) []int {
	ans := make([]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			ans = append(ans, i)
		}
	}
	return ans
}
