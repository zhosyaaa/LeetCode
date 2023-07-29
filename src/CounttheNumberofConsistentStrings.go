package main

func countConsistentStrings(allowed string, words []string) int {
	table := make(map[rune]bool)
	for i := 0; i < len(allowed); i++ {
		table[rune(allowed[i])] = true
	}
	count := 0
	for i := 0; i < len(words); i++ {
		temp := words[i]
		b := true
		for j := 0; j < len(temp); j++ {
			if table[rune(temp[j])] == true {
				b = true
			} else {
				b = false
				break
			}
		}
		if b == true {
			count++
		}
	}
	return count
}
