package main

func sortArrayByParity(nums []int) []int {
	even := make([]int,  0)
	odd := make([]int, 0)
	for _, val := range nums{
		if isEven(val){
			even = append(even, val)
			continue
		}
		odd = append(odd, val)
	}
	res := make([]int, 0)
	copy(res, even)
	copy(res, odd)
	return res
}
func isEven(num int) bool{
	if num%2==0{
		return true
	}
	return false
}

func sortArrayByParityII(nums []int) []int {
	even := make([]int,  0)
	odd := make([]int, 0)
	for _, val := range nums{
		if isEven(val){
			even = append(even, val)
			continue
		}
		odd = append(odd, val)
	}
	res := make([]int, 0)
	k, j := 0, 0
	for i := range nums{
		if isEven(i){
			res = append(res, even[k])
			k++
			continue
		}
		res = append(res, odd[j])
		j++
	}
	return res
}