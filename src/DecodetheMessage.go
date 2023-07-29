package main

func decodeMessage(key string, message string) string {
	table := make(map[rune]rune)
	for i := 0; i < len(key); i++ {
		if _, ok := table[rune(key[i])]; key[i] != ' ' && !ok {
			table[rune(key[i])] = rune(int('a') + len(table))
		}
	}
	res := make([]rune, 0)
	for i := 0; i < len(message); i++ {
		if message[i] == ' ' {
			res = append(res, rune(message[i]))
			continue
		}
		res = append(res, table[rune(message[i])])
	}
	return string(res)
}
