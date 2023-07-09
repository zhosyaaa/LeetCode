package main

//	func main() {
//		arr := []int{1, 2, 3, 1, 2, 3}
//		fmt.Println(containsNearbyDuplicate(arr, 2))
//	}
func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if j-i > k {
				continue
			}
			if nums[i] == nums[j] && j <= k+i {
				return true
			}
		}
	}
	return false
}
