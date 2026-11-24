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

// -------------------------------------------------

func addDigits2(num int) int {
	flag := true
	result := 0
	for flag {
		sum := 0
		for num > 0 {
			temp := num % 10
			sum += temp
			num /= 10
		}
		if sum < 10 {
			flag = false
		} else {
			num = sum
		}
		result = sum
	}
	return result
}
