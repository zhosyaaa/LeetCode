package main

//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//	tArr, sArr := strings.Split(t, ""), strings.Split(s, "")
//	sort.Strings(tArr)
//	sort.Strings(sArr)
//	for i, val := range tArr {
//		if val != sArr[i] {
//			return false
//		}
//	}
//	return true
//}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := make(map[rune]int)
	for _, val := range t {
		counts[val]++
	}
	for _, val := range s {
		counts[val]--
	}
	for _, val := range t {
		if counts[val] != 0 {
			return false
		}
	}
	return true
}
