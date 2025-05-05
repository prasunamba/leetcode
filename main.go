package main

import (
	"fmt"
)

func main() {

	// nums := []int{4, 2, 5, 7, 3}
	for i := range 10 {

		result := fib(i)
		fmt.Print(" ", result)
	}
}
func fib(num int) int {
	if num <= 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return fib(num-1) + fib(num-2)
	}
}
