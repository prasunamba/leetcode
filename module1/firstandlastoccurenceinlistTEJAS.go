package module1

import "fmt"

func Tejasquestion() {
	list := []int{8, 5, 7, 7, 8, 8, 10, 7, 7, 8, 8, 10}
	num := 8
	var count []int
	for index, n := range list {
		if num == n {
			count = append(count, index)
		}
	}
	// Print results safely
	if len(count) > 0 {
		fmt.Println("Indices:", count)
		fmt.Println("First occurrence:", count[0])
		fmt.Println("Last occurrence:", count[len(count)-1])
	} else {
		fmt.Println(num, "not found in list")
	}
}
