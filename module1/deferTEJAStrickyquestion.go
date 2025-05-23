package module1

import "fmt"

func Trickydefer() int {
	number := 0
	if number == 0 {
		return 0
	}
	defer func() {
		fmt.Println("won't  execute ")
	}()
	if number == 1 {
		return 1
	}
	return number
}

/* as func is returning before the encounter of the defer statement
defer won't be executed , to execute the defer statement the defer statement should register itself
before the return of the function */
