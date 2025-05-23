package module1

import "fmt"

func Add(a, b int) int {
	sum := a + b
	fmt.Printf("sum of %d and %d is %d \n", a, b, sum)
	return sum
}
