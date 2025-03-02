package module1

import (
	"fmt"
	"strings"
)

func CharFrequency() {
	name := "gnanaprasunamba"
	name = strings.ReplaceAll(name, " ", "")
	runemap := make(map[rune]int)
	for _, char := range name {
		runemap[char]++
	}
	for char, count := range runemap {
		fmt.Printf("'%c' : %d\n", char, count)
	}
}
func CharFrequencyOptimized() {
	name := "gnana prasunamba"
	name = strings.ReplaceAll(name, " ", "")
	fmt.Println("name", name)
	map1 := make(map[rune]int)
	runes := []rune(name)
	for i, j := 0, len(runes)-1; i <= j; i, j = i+1, j-1 {
		map1[runes[i]]++
		if i != j { // Avoid double-counting if i == j (for odd-length strings)
			map1[runes[j]]++
		}
	}
	for key, value := range map1 {
		fmt.Printf("'%c' : %d\n", key, value)
	}
}
