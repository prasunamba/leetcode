package main

import "fmt"

func reverseslice() {

	a := []int{1, 2, 3, 4, 5}
	b := make([]int, len(a))
	fmt.Println("before a ", a)
	j := len(a)
	for index, i := range a {
		b[j-1-index] = i
	}
	fmt.Println("after b  ", b)
}

func main() {
	reverseslice()
}
