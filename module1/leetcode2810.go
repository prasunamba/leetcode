package module1

import "fmt"

func FinalString(s string) string {
	result := ""
	for _, value := range s {
		if value == 'i' {
			sname := []rune(result)
			result = ""
			fmt.Println("res", result)
			for i, j := 0, len(sname)-1; i < j; i, j = i+1, j-1 {
				sname[i], sname[j] = sname[j], sname[i]
			}
			fmt.Println("snamae", sname)
			fmt.Println("sname", string(sname))
			result += string(sname)

		} else {
			result += string(value)
		}
	}
	return result
}
