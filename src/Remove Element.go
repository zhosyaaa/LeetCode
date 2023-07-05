package main

//	func main() {
//		arr:= []int{1, 2, 5, 4, 8}
//		val := 2
//		fmt.Println(removeElement(arr, val))
//	}
func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	return j
}
