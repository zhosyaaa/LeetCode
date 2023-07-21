package main

import "math"

//func main() {
//	println(reverse(1534236469))
//}

func reverse(x int) int {
	res := 0
	for x != 0 {
		res = (res * 10) + x%10
		x = x / 10
	}
	if res > int(math.Pow(2, 31)-1) || res < -int(math.Pow(2, 31)) {
		return 0
	}
	return res
}
