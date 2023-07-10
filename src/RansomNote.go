package main

import (
	"fmt"
	"strings"
)

//func main() {
//	fmt.Println(canConstruct("aa", "ab"))
//}

func canConstruct(ransomNote string, magazine string) bool {
	words := make(map[string]int)
	arr := strings.Split(magazine, "")
	for i := 0; i < len(arr); i++ {
		words[arr[i]] += 1
	}
	fmt.Println(words)
	for _, val := range ransomNote {
		if words[string(val)] == 0 {
			return false
		}
		words[string(val)] -= 1
	}
	return true
}
