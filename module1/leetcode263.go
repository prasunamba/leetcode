package module1

import "fmt"

func IsUgly(n int) bool {
	if n <= 0 {
		return false
	}

	// Keep dividing n by 2, 3, and 5
	for n%2 == 0 {
		n /= 2
	}
	fmt.Println("", n)
	for n%3 == 0 {
		n /= 3
	}
	fmt.Println("", n)
	for n%5 == 0 {
		n /= 5
	}

	// If n is reduced to 1, it's an ugly number
	return n == 1
}
