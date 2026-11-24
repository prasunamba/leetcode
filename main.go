package main

import (
	"fmt"
)

func totalMoney(n int) int {
	total := 0
	monday := 1

	for n > 0 {
		deposit := monday

		// Process one week (maximum 7 days)
		for i := 0; i < 7 && n > 0; i++ {
			total += deposit
			deposit++
			n--
		}

		monday++
	}

	return total
}

func main() {
	fmt.Println("", totalMoney(20))
}
