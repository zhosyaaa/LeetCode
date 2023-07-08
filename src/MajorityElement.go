package main

//func main() {
//	arr := []int{3, 3, 4}
//	fmt.Println(majorityElement(arr))
//}

func majorityElement(nums []int) int {
	max, maxCount, count := nums[0], 0, 0
	for i := 0; i < len(nums); i++ {
		count = 0
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
			max = nums[i]
		}

	}
	return max
}
