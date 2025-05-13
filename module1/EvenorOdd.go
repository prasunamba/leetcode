package module1

import "fmt"

func EvenorOdd(num int) bool {
	if num%2 == 0 {
		fmt.Printf("%d is Even ", num)
		return true
	} else {
		fmt.Printf("%d is Odd ", num)
		return false
	}
}
