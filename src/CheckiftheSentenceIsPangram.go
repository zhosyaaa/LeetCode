package main

func checkIfPangram(sentence string) bool {
	table := make(map[rune]int)
	for i := 0; i < len(sentence); i++ {
		table[rune(sentence[i])] += 1
	}
	if len(table) < 26 {
		return false
	}
	return true
}
