package main

//func main() {
//	arr := []int{4, 5, 6, 7, 0, 1, 2}
//	fmt.Println(findMin(arr))
//}

func findMin(nums []int) int {
	if len(nums) == 3 && nums[1] < nums[0] && nums[1] < nums[2] {
		return nums[1]
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		if nums[left] <= nums[right] {
			return nums[left]
		}
		middle := (left + right) / 2
		if nums[left] > nums[middle] {
			right = middle
		} else if nums[middle] > nums[right] {
			left = middle + 1
		}
	}
	if nums[left] <= nums[right] {
		return nums[left]
	}
	return -1
}
