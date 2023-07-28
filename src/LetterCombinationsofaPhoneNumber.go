package main

func letterCombinations(digits string) []string {
	combination := make(map[string][]string)
	combination["2"] = []string{"a", "b", "c"}
	combination["3"] = []string{"d", "e", "f"}
	combination["4"] = []string{"g", "h", "i"}
	combination["5"] = []string{"j", "k", "l"}
	combination["6"] = []string{"m", "n", "o"}
	combination["7"] = []string{"p", "q", "r", "s"}
	combination["8"] = []string{"t", "u", "v"}
	combination["9"] = []string{"w", "x", "y", "z"}
	var mappedValue []string
	if len(digits) > 0 {
		mappedValue = combination[string(digits[0])]
	}
	if len(digits) <= 1 {
		return mappedValue
	}
	return makeCombinations(mappedValue, letterCombinations(digits[1:]))
}

func makeCombinations(first, second []string) []string {
	resultant := make([]string, 0, len(first)*len(second))
	for i := 0; i < len(first); i++ {
		for j := 0; j < len(second); j++ {
			resultant = append(resultant, first[i]+second[j])
		}
	}
	return resultant
}
