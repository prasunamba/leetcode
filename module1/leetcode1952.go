package module1

import (
	"fmt"
	"math"
)

func IsThree(n int) bool {
	nn := math.Sqrt(float64(n))
	result := 0
	fmt.Println("nn", nn)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result++
		}
	}
	fmt.Println("nn", nn, result)
	if result == 3 {
		return true
	}
	return false
}
