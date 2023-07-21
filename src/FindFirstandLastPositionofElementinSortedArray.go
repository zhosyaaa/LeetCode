package main

//func main() {
//	arr := []int{5, 8, 7, 8, 8, 10}
//	fmt.Println(searchRange(arr, 8))
//}

func searchRange(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{-1, -1}
	}
	left := 0
	right := len(nums) - 1
	resLeft := -1
	resRight := -1

	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}

	if nums[left] == target {
		resLeft = left
	} else if nums[right] == target {
		resLeft = right
	} else {
		return []int{-1, -1}
	}

	left = 0
	right = len(nums) - 1
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid
		} else {
			left = mid
		}
	}

	if nums[right] == target {
		resRight = right
	} else {
		resRight = left
	}

	return []int{resLeft, resRight}
}
