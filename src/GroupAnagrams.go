package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	isPresent := make(map[string][]string)
	for _, val := range strs {
		sortVal := sortStr(val)
		isPresent[sortVal] = append(isPresent[sortVal], val)
	}
	res := make([][]string, 0, len(isPresent))
	for _, val := range isPresent {
		res = append(res, val)
	}
	return res
}

func sortStr(s string) string {
	sorted := strings.Split(s, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}
