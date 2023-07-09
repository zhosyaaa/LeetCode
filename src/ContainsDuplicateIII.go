package main

//func main() {
//	arr := []int{8, 7, 15, 1, 6, 1, 9, 15}
//	fmt.Println(containsNearbyAlmostDuplicate(arr, 1, 3))
//}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j <= i+indexDiff && j < len(nums); j++ {
			if absInt(nums[i], nums[j]) <= valueDiff {
				return true
			}
		}
	}
	return false
}

func absInt(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
