package module1

func ClearDigits(s string) string {
	runeslice := []rune(s)
	var newruneslice []rune

	for _, value := range runeslice {
		if value >= '0' && value <= '9' {
			if len(newruneslice) > 0 {
				newruneslice = newruneslice[:len(newruneslice)-1]
			}
		} else {
			newruneslice = append(newruneslice, value)
		}
	}

	return string(newruneslice)
}
