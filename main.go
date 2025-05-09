package main

import (
	"fmt"
)

func main() {

	nums := []int{4, 2, 5, 7, 3}
	// nums := os.Args[1]
	// num, _ := strconv.Atoi(nums)
	for i := range nums {

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
