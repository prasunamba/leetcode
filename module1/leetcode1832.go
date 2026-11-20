package module1

func checkIfPangram(sentence string) bool {
	panmap := make(map[rune]int)
	for _, value := range sentence {
		panmap[value]++
	}
	return len(panmap) == 26
}
