package main

func getConcatenation(nums []int) []int {
	newLen := 2 * len(nums)
	ans := make([]int, 0, newLen)
	for i := 0; i <= newLen/2 && len(ans) != newLen; i++ {
		if i == newLen/2 {
			i = 0
		}
		ans = append(ans, nums[i])
	}
	return ans
}
