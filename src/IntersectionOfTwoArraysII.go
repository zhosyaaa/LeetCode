package main

//	func main() {
//		arr1 := []int{1, 2, 5, 4, 8}
//		arr2 := []int{4, 2}
//		fmt.Println(intersection(arr1, arr2))
//	}
func intersect(nums1 []int, nums2 []int) []int {
	var count = map[int]int{}
	var result = []int{}
	for _, num := range nums1 {
		count[num] += 1
	}
	for _, num := range nums2 {
		if count[num] != 0 {
			result = append(result, num)
			count[num] -= 1
		}
	}
	return result

}
