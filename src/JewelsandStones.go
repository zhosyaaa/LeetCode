package main

import (
	"strings"
)

func numJewelsInStones(jewels string, stones string) int {
	typeJewels := strings.Split(jewels, "")
	mapJewels := make(map[string]int)
	for i := 0; i < len(typeJewels); i++ {
		mapJewels[typeJewels[i]] += 1
	}
	count := 0
	for i := 0; i < len(stones); i++ {
		if mapJewels[string(stones[i])] != 0 {
			count++
		}
	}
	return count
}
