package main

func romanToInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'M' {
			res += 1000
		}
		if s[i] == 'D' {
			res += 500
		}
		if s[i] == 'C' {
			if i+1 < len(s) && s[i+1] == 'D' {
				res += 400
				i++
				continue
			}
			if i+1 < len(s) && s[i+1] == 'M' {
				res += 900
				i++
				continue
			}
			res += 100
		}
		if s[i] == 'L' {
			res += 50
		}
		if s[i] == 'X' {
			if i+1 < len(s) && s[i+1] == 'L' {
				res += 40
				i++
				continue
			}
			if i+1 < len(s) && s[i+1] == 'C' {
				res += 90
				i++
				continue
			}
			res += 10
		}
		if s[i] == 'V' {
			res += 5
		}
		if s[i] == 'I' {
			if i+1 < len(s) && s[i+1] == 'V' {
				res += 4
				i++
				continue
			}
			if i+1 < len(s) && s[i+1] == 'X' {
				res += 9
				i++
				continue
			}
			res++
		}
	}
	return res
}
