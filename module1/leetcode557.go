pcakage module1

func ReverseWords(s string) string {
	words := strings.Fields(s)
	result := []string{}
	fmt.Println("", words)
	for _, s := range words {
		ss := []rune(s)
		for i, j := 0, len(ss)-1; i <= j; i, j = i+1, j-1 {
			ss[i], ss[j] = ss[j], ss[i]
		}
		result = append(result, string(ss))
		fmt.Println("s,ss", s, ss)
	}
	fmt.Println("result", result)
	return strings.Join(result, " ")
}