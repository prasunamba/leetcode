package module1

import "fmt"

/*
f(0)=0
f(1)=1
f(num-1)+f(num-2)
*/
func FibonacciSeries(num int) int {
	if num <= 0 {
		return 0
	} else if num == 1 {
		return 1
	}
	return FibonacciSeries(num-1) + FibonacciSeries(num-2)

}
func Fib(num int) {
	for i := range num {
		fmt.Printf("%d ", FibonacciSeries(i))
	}
}
