pcakage module1

func ReverseStr(s string, k int) string {
	result := ""
	if len(s) == 1 {
		return s
	}
	for i := 0; i < len(s); i = i + 2*k {
		end := i + 2*k
		if end > len(s) {
			end = len(s)
		}
		fmt.Println("", s[i:end])
		l := len(s[i:end])

		strslice := []rune(s[i:end])
		if l == 1 {
			result += string(strslice)
		} else if l < k {
			// reverse all
			for j, kk := 0, l-1; j <= kk; j, kk = j+1, kk-1 {
				strslice[j], strslice[kk] = strslice[kk], strslice[j]
			}
			fmt.Println("l<k", string(strslice))
			result += string(strslice)
		} else {
			// reverse the first k
			for j, p := 0, k-1; j <= p; j, p = j+1, p-1 {
				strslice[j], strslice[p] = strslice[p], strslice[j]
			}

			fmt.Println("else", string(strslice))
			result += string(strslice)
		}
	}
	fmt.Println("result", result)
	return result
}