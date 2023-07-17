package main

//func main() {
//	arr := []int{2, 2, 2, 0, 1}
//	fmt.Println(findMinII(arr))
//}

func findMinII(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		middle := left + (right-left)/2
		if nums[left] > nums[middle] {
			right = middle
			left++
		} else if nums[middle] > nums[right] {
			left = middle + 1
		} else {
			right--
		}
	}
	return nums[left]
}
