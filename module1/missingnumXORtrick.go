package module1

import "fmt"

func MissingnumberXOR() {
	nums := []int{3, 0, 1}
	missing := len(nums)
	for i, num := range nums {
		missing ^= i ^ num
	}
	fmt.Println("", missing)
}
