package main



func numTilePossibilities(tiles string) int {
	chars := make([]int, 26)
	for _, val := range tiles {
		chars[val-'A'] += 1
	}
	return dfs(chars)
}

func dfs(chars []int) int {
	sum := 0
	for i := 0; i < 26; i++ {
		if chars[i] == 0 {
			continue
		}
		sum++
		chars[i]--
		sum += dfs(chars)
		chars[i]++
	}
	return sum
}