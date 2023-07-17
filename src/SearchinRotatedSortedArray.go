package main

//	func main() {
//		arr := []int{4, 5, 6, 7, 0, 1, 2}
//		fmt.Println(search(arr, 6))
//	}
func search(nums []int, target int) int {
	sec, ind := sorted(nums)
	res := binary(sec, target)
	if res == -1 {
		return binary(nums[:ind], target)
	}
	return res + ind
}
func binary(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] < target {
			left = middle + 1
		}
		if nums[middle] > target {
			right = middle - 1
		}
	}
	return -1
}

func sorted(nums []int) ([]int, int) {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return nums[i+1:], i + 1
		}
	}
	return nums, 0
}
