package main

//	func main() {
//		arr := []int{1, 2, 5, 4, 8}
//		fmt.Println(removeDuplicates(arr))
//	}
func removeDuplicates(nums []int) int {
	ln := len(nums)
	if ln <= 1 {
		return ln
	}
	j := 0
	for i := 0; i < ln; i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}
