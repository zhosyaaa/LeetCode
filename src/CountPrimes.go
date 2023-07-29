package main

import "math"

func countPrimes(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	a := math.Sqrt(float64(n))
	for i := 2; i <= int(a); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
