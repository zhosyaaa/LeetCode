package main

func buildArray(nums []int) []int {
	ans := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		ans = append(ans, nums[nums[i]])
	}
	return ans
}
