package main

func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if num == target || num > target {
			return i
		}
	}
	return len(nums)
}
