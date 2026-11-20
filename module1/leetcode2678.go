package module1

import (
	"fmt"
	"strconv"
)

func CountSeniors(details []string) int {
	count := 0
	for _, value := range details {
		num, _ := strconv.Atoi(value[11:13])
		fmt.Println("l", value[11:13], count, num)
		if num > 60 {
			count++
		}
	}
	return count
}
