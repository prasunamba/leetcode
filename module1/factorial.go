package module1

import (
	"fmt"
)

func fact(i int) int {
	if i < 2 {
		return 1
	}
	return fact(i-1) * i
}
func Factorial(num int) int {

	result := fact(num)
	fmt.Println("result", result)
	return result
}
