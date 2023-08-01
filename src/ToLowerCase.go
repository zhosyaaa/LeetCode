package main

import (
	"sort"
	"strings"
)

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

// Sort the People
func sortPeople(names []string, heights []int) []string {
	sorted := make(map[int]string)
	res := make([]string, 0)
	for i, val := range heights {
		sorted[val] = names[i]
	}
	sort.Ints(heights)
	for i := len(heights) - 1; i >= 0; i-- {
		res = append(res, sorted[heights[i]])
	}
	return res
}

// Truncate Sentence
func truncateSentence(s string, k int) string {
	ans := strings.Split(s, " ")
	return strings.Join(ans[:k], " ")
}
