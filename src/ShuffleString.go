package main

import "strings"

func restoreString(s string, indices []int) string {
	index := make(map[int]rune, len(indices))
	for i := 0; i < len(indices); i++ {
		index[indices[i]] = rune(s[i])
	}
	res := make([]string, 0)
	for i := 0; i < len(indices); i++ {
		res = append(res, string(index[i]))
	}
	return strings.Join(res, "")
}
