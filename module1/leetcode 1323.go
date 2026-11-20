package module1

import "fmt"

func Maximum69Number(num int) int {
	var numarr []int
	var result int
	flag := true
	for num > 0 {
		temp := num % 10
		numarr = append(numarr, temp)
		fmt.Println("", numarr)
		num = num / 10
	}
	for i := len(numarr) - 1; i >= 0; i-- {
		fmt.Println("hello", result, numarr[i])

		if numarr[i] == 6 && flag {
			numarr[i] = 9
			flag = false
		}
		result = result*10 + numarr[i]

	}

	return result
}
