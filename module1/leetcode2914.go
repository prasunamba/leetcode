package module1

func MinChanges(s string) int {
	count := 0
	for i := 0; i < len(s); i = i + 2 {
		if s[i] != s[i+1] {
			count++
		}
	}
	return count
}
