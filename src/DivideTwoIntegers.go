package main

import (
	"math"
)

//	func main() {
//		fmt.Println(divide(7, -3))
//	}
func divide(dividend int, divisor int) int {
	min := -int(math.Pow(2, 31))
	max := int(math.Pow(2, 31) - 1)
	if dividend/divisor > max {
		return max
	}
	if dividend/divisor < min {
		return min
	}
	return dividend / divisor
}
