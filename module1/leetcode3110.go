package module1

import (
	"fmt"
	"main/utils"
)

func ScoreOfString(s string) int {
	sum := 0

	for i := 0; i < len(s)-1; i++ {
		a := int(s[i])
		b := int(s[i+1])

		if a > b {
			sum += a - b
		} else {
			sum += b - a
		}
	}

	return sum
}

func ScoreOfString2(s string) int {

	sum := 0
	runearr := []rune(s)

	for i := 0; i < len(runearr)-1; i++ {
		num := utils.Abs(int(runearr[i] - runearr[i+1]))
		fmt.Print(" ", num)
		sum += num
	}
	return sum
}
