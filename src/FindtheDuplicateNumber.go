package main


func find

Duplicate(nums []int) int {
	n := len(nums)-1
	counts := make(map[int]int)
	for i := 0; i < n+1; i++ {
		counts[nums[i]] += 1
		if counts[nums[i]] > 1 {
			return nums[i]
		}
	}
	return 0
}