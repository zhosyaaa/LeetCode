package main

import (
	"strings"
)

func sortSentence(s string) string {
	words := strings.Split(s, " ")
	ans := make([]string, len(words))
	for _, val := range words {
		lenth := len(val)
		ind := val[lenth-1] - 48
		word := val[:lenth-1]
		ans[ind-1] = word
	}
	return strings.Join(ans, " ")
}
