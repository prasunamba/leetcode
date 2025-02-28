package module1

import "fmt"

func MinMax() {
	slice := []int{2, 1, 3, 55, 22}
	min, max := slice[0], slice[0]
	for _, num := range slice {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	fmt.Println("", min, max)
}
