package module1

import "fmt"

func Reversenumber() {
	num := 123
	x := num
	var r, temp int
	for num > 0 {
		r = num % 10
		temp = temp*10 + r
		num = num / 10
	}
	fmt.Printf("before : %d after : %d\n ", x, temp)
}
