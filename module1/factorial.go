package module1

import (
	"fmt"
)

/*
f(0)=1
f(num-1)*num
*/
func fact(i int) int {
	if i < 2 {
		return 1
	}
	return fact(i-1) * i
}
func Factorial(num int) int {

	result := fact(num)
	fmt.Printf("factorial of %d is %d\n ", num, result)
	return result
}
