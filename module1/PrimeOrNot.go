package module1

import "math"

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
