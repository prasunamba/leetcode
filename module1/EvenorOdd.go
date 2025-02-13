package module1

import "fmt"

func EvenorOdd(num int) {
	if num%2 == 0 {
		fmt.Printf("%d is Even ", num)
	} else {
		fmt.Printf("%d is Odd ", num)
	}
}
