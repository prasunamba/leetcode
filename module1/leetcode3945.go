package module1

import "fmt"

func DigitFrequencyScore(n int) int {
	digitmap := make(map[int]int)
	for n > 0 {
		temp := n % 10
		digitmap[temp]++
		n = n / 10
	}
	fmt.Println("digitmap", digitmap)
	sum := 0
	for key, value := range digitmap {
		sum += key * value
	}
	return sum
}
