package main

//func main() {
//	fmt.Println(getRow(5))
//}

func getRow(rowIndex int) []int {
	ans := [][]int{[]int{1}}
	res := ([]int{})
	for i := 0; i < rowIndex+1; i++ {
		v := make([]int, i+1)
		v[0] = 1
		v[i] = 1
		for j := 1; j < len(v)-1; j++ {
			v[j] = ans[i][j] + ans[i][j-1]
		}
		ans = append(ans, v)
		if i == rowIndex {
			res = v
		}
	}
	return res
}
