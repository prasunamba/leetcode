package module1

import (
	"fmt"
	"math"
)

func PrimeOrNot(num int) bool {
	prime := true

	// Edge cases: 0 and 1 are not prime
	if num < 2 {
		return false
	}
	if num == 2 {
		return true // 2 is a prime number
	}

	// Check divisibility from 2 to sqrt(num)
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ { // Fixed loop condition
		if num%i == 0 {
			prime = false
			break
		}
	}
	return prime
}
func Closestprimenumbersinrange() {
	left := 69346
	right := 69379
	var primeslice []int
	for i := left; i <= right; i++ {
		if PrimeOrNot(i) {
			primeslice = append(primeslice, i)
		}
	}
	mindiff := math.MaxInt32
	var result []int
	fmt.Println("primeslice", primeslice)
	if len(primeslice) >= 2 {

		for index := 1; index < len(primeslice); index++ {
			if primeslice[index]-primeslice[index-1] < mindiff {
				mindiff = primeslice[index] - primeslice[index-1]
				result = []int{primeslice[index-1], primeslice[index]}
			}
		}
	} else {
		result = []int{-1, -1}
		fmt.Println("", result)
	}
	fmt.Println("", result)
}
