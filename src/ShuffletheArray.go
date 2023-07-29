package main

func shuffle(nums []int, n int) []int {
	ans := make([]int, 0, 2*n)
	left, right := make([]int, 0, n), make([]int, 0, n)
	left = append(left, nums[:n]...)
	right = append(right, nums[n:]...)
	ans = append(ans, left[0])
	a, b := 1, 0
	for i := 1; i < len(nums); i++ {
		if i%2 != 0 {
			ans = append(ans, right[b])
			b++
			continue
		}
		ans = append(ans, left[a])
		a++

	}
	return ans
}
