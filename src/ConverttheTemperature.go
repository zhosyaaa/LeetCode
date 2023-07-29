package main

func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.80 + 32.00}
}

// Add Two Integers
func sum(num1 int, num2 int) int {
	return num1 + num2
}

// Smallest Even Multiple
func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return 2 * n
}

// Subtract the Product and Sum of Digits of an Integer
func subtractProductAndSum(n int) int {
	num := 0
	nums := make([]int, 0)
	for n != 0 {
		num = n % 10
		n /= 10
		nums = append(nums, num)
	}
	product := 1
	sum := 0
	for i := 0; i < len(nums); i++ {
		product *= nums[i]
		sum += nums[i]
	}
	return product - sum
}

// Sum Multiples
func sumOfMultiples(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			res += i
		}
	}
	return res
}

// Count the Digits That Divide a Number
func countDigits(num int) int {
	number := 0
	nums := make([]int, 0)
	temp := num
	for num != 0 {
		number = num % 10
		num /= 10
		nums = append(nums, number)
	}
	count := 0
	for _, val := range nums {
		if temp%val == 0 {
			count++
		}
	}
	return count
}
