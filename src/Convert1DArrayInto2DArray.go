package main

func construct2DArray(original []int, m int, n int) [][]int {
	var result [][]int
	if len(original) != m*n {
		return [][]int{}
	}
	for i := 0; i<n*m; i+=n {
		row := original[i:i+n]
		result = append(result, row)
	}
	return result
}