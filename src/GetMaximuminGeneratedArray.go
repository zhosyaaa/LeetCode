package main

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	nums := make([]int, n+1, n+1)
	nums[0] = 0
	nums[1] = 1
	for i := 0; i <= n; i++ {
		if 2 <= 2*i && 2*i <= n {
			nums[2*i] = nums[i]
		}
		if 2 <= 2*i+1 && 2*i+1 <= n {
			nums[2*i+1] = nums[i] + nums[i+1]
		}
	}
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
