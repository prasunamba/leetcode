package module1
package main

import "fmt"

/* func main() {
	num := make(chan int)
	go run(num)
	ok := <-num
	fmt.Println("ok", ok)
} */
func run(num chan int) {
	for i := 0; i < 3; i++ {
		fmt.Println("hello")
	}
	num <- 10
	close(num)
}
