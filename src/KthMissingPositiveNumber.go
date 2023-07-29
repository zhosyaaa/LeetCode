package main

func findKthPositive(arr []int, k int) int {
	num, res := 1, 0
	for k != 0 {
		if !isPresent(arr, num) {
			res = num
			k--
		}
		num++
	}
	return res
}

func isPresent(arr []int, k int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == k {
			return true
		}
	}
	return false
}
