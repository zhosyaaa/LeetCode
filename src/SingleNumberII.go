package main

//
//func main() {
//	arr := []int{0, 1, 0, 1, 0, 1, 99}
//	fmt.Println(singleNumber(arr))
//}

func singleNumber(nums []int) int {
	counts := make(map[int]int)
	for _, val := range nums {
		counts[val] += 1
	}
	for i, val := range counts {
		if val == 1 {
			return i
		}
	}
	return 0
}
