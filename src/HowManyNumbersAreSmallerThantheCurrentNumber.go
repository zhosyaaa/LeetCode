package main

func smallerNumbersThanCurrent(nums []int) []int {
	ans := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		ans = append(ans, countSmaller(nums[i], nums))
	}
	return ans
}

func countSmaller(num int, nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if num > nums[i] {
			count++
		}
	}
	return count
}
