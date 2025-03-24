package module1

import "fmt"

func helpfunc(num int) int {
	temp := 0

	for num > 0 {
		r := num % 10
		temp += r
		num = num / 10
	}
	fmt.Println("temp", temp)
	return temp
}
func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	for !(0 < num && num <= 9) {
		num = helpfunc(num)
	}
	return num
}
