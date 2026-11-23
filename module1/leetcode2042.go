package module1

import (
	"fmt"
	"strconv"
	"strings"
)

func AreNumbersAscending(s string) bool {
	words := strings.Fields(s)
	fmt.Println("", words)
	maxnum := 0
	for _, value := range words {
		num, _ := strconv.Atoi(value)
		fmt.Println("", value, num)
		if num != 0 && num > maxnum {
			maxnum = num
		} else if num != 0 {
			return false
		}

	}
	return true
}
