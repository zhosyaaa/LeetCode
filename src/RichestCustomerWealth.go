package main

func maximumWealth(accounts [][]int) int {
	max, sum := 0, 0
	for i := 0; i < len(accounts); i++ {
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		if max < sum {
			max = sum
		}
		sum = 0
	}
	return max
}
