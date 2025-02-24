package module1

import "fmt"

func CharFrequency() {
	name := "gnanaprasunamba"
	runemap := make(map[rune]int)
	for _, char := range name {
		runemap[char]++
	}
	for char, count := range runemap {
		fmt.Printf("'%c' : %d\n", char, count)
	}
}
