package module1

import "fmt"

func test() int {
	number := 0
	if number == 0 {
		return 0
	}
	defer func() {
		fmt.Println("exit")
	}()
	if number == 1 {
		return 1
	}
	return number
}
