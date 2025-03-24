package module1

import "fmt"

func help(num int) int {
	temp := 0

	for num > 0 {
		r := num % 10
		temp += r * r
		num = num / 10
	}
	fmt.Println("temp", temp)
	return temp
}
func happy(num int) bool {
	map1 := make(map[int]bool)
	for num != 1 && !map1[num] {
		map1[num] = true
		num = help(num)
	}
	return num == 1
}
