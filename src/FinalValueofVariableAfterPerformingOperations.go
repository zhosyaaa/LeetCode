package main

import "strings"

func finalValueAfterOperations(operations []string) int {
	res := 0
	for i := 0; i < len(operations); i++ {
		if operations[i] == "++X" || operations[i] == "X++" {
			res++
		}
		if operations[i] == "--X" || operations[i] == "X--" {
			res--
		}
	}
	return res
}

// Goal Parser Interpretation
func interpret(command string) string {
	res := make([]string, 0)
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			res = append(res, "G")
		}
		if command[i] == '(' {
			if command[i+1] == ')' {
				res = append(res, "o")
				i++
			} else {
				res = append(res, "al")
				i += 3
			}
		}
	}
	return strings.Join(res, "")
}

// Maximum Number of Words Found in Sentences
func mostWordsFound(sentences []string) int {
	max := 0
	for i := 0; i < len(sentences); i++ {
		arr := make([]string, 0)
		arr = append(arr, strings.Split(sentences[i], " ")...)
		len := len(arr)
		if max < len {
			max = len
		}
	}
	return max
}
