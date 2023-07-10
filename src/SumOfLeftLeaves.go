package main

//	func main() {
//		fmt.Println(countSegments("aaa     aa"))
//	}
func countSegments(s string) int {
	isPresent := true
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			isPresent = true
		}
		if s[i] != ' ' && isPresent {
			count++
			isPresent = false
		}
	}
	return count
}
