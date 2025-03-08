package main

import (
	"fmt"
)

func main() {

}

func Matrixadd() {
	dynmat := make([][]int, 3) // 3 rows
	for i := range dynmat {
		dynmat[i] = make([]int, 3) // 3 col for each row
	}
	fmt.Println("", dynmat)
	for i := range dynmat {
		num1 := []int{1, 2, 3}
		dynmat[i] = num1

	}
	fmt.Println("", dynmat)
}
