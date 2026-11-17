package module1

import (
	"fmt"
	"strconv"
)

func DigitSum(s string, k int) string {
	for len(s) > k {
		news := ""
		fmt.Println(len(s), s)
		for i := 0; i < len(s); i = i + k {
			end := i + k
			if end > len(s) {
				end = len(s)
			}
			fmt.Println(s[i:end])
			sum := 0
			for _, ch := range s[i:end] {
				sum += int(ch - '0') // because int cant take the more then 20 digits
			}

			news += strconv.Itoa(sum)
			fmt.Println("news", news)
		}
		s = news
	}
	return s
}
