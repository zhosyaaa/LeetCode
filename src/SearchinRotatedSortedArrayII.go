package main

func searchII(nums []int, target int) bool {
	sec, ind := sorted(nums)
	res := binaryBool(sec, target)
	if !res {
		return binaryBool(nums[:ind], target)
	}
	return res
}

func binaryBool(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return true
		}
		if nums[middle] < target {
			left = middle + 1
		}
		if nums[middle] > target {
			right = middle - 1
		}
	}
	return false
}
