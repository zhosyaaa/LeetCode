package main


// taks 1 Majority Element -easy
//func majorityElement(nums []int) int {
//	max, maxCount, count := nums[0], 0, 0
//	for i := 0; i < len(nums); i++ {
//		count = 0
//		for j := 0; j < len(nums); j++ {
//			if nums[i] == nums[j] {
//				count++
//			}
//		}
//		if count > maxCount {
//			maxCount = count
//			max = nums[i]
//		}
//
//	}
//	return max
//}

// /task 2 Is Subsequence -easy
func isSubsequence(s string, t string) bool {
	res := make([]int, 0, len(s))
	arrS, arrT := []rune(s), []rune(t)
	if len(arrS) == 0 {
		return true
	}
	lastInd := 0
	for i := 0; i < len(arrS); i++ {
		ind, ok := findChar(arrS[i], arrT)
		if !ok {
			return false
		}
		if ind != -1 {
			if len(res) != 0 {
				lastInd = res[len(res)-1]
			}
			res = append(res, ind+lastInd)
		}
		arrT = arrT[ind+1:]
	}
	if len(res) == 0 {
		return false
	}
	for i := 0; i < len(res)-1; i++ {
		if res[i] > res[i+1] {
			return false
		}
	}
	return true
}
func findChar(char rune, t []rune) (int, bool) {
	ind := -1
	for i := 0; i < len(t); i++ {
		if t[i] == char {
			ind = i
			return ind, true
		}
	}
	return ind, false
}

// task 4 167. Two Sum II - Input Array Is Sorted medium
func twoSum(numbers []int, target int) []int {
	res := make([]int, 2)
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if target == sum {
			res[0] = left + 1
			res[1] = right + 1
			return res
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return res
}

//// task 6 Search in Rotated Sorted Array Medium
//func search(nums []int, target int) int {
//	sec, ind := sorted(nums)
//	res := binary(sec, target)
//	if res == -1 {
//		return binary(nums[:ind], target)
//	}
//	return res + ind
//}
//func binary(nums []int, target int) int {
//	left := 0
//	right := len(nums) - 1
//	for left <= right {
//		middle := (left + right) / 2
//		if nums[middle] == target {
//			return middle
//		}
//		if nums[middle] < target {
//			left = middle + 1
//		}
//		if nums[middle] > target {
//			right = middle - 1
//		}
//	}
//	return -1
//}
//
//func sorted(nums []int) ([]int, int) {
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i] > nums[i+1] {
//			return nums[i+1:], i + 1
//		}
//	}
//	return nums, 0
//}
