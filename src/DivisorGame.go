package main

func divisorGame(n int) bool {
	count := 0
	for {
		n = divisor(n)
		if n == 0 {
			break
		}
		count++
	}
	if count%2 == 0 {
		return false
	}
	return true
}

func divisor(n int) int {
	for i := 1; i < n; i++ {
		if n%i == 0 {
			return n - i
		}
	}
	return 0
}
