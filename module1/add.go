package module1

import "fmt"

func Add(a, b int) int {
	fmt.Printf("sum of %d  and %d is %d \n ", a, b, a+b)
	return a + b
}
