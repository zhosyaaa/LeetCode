package main

func missingNumber(nums []int) int {
	len := len(nums)
	totalSum := len * (len + 1) / 2
	for _, val := range nums {
		totalSum -= val
	}
	return totalSum
}
