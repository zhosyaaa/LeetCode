package main

//func main() {
//	fmt.Println(generate(11))
//}

func generate(rowIndex int) [][]int {
	ans := [][]int{[]int{1}}

	for i := 0; i < rowIndex; i++ {
		v := make([]int, i+1)
		v[0] = 1
		v[i] = 1
		for j := 1; j < len(v)-1; j++ {
			v[j] = ans[i][j] + ans[i][j-1]
		}
		ans = append(ans, v)
	}
	return ans[1:]
}
