package module1

import (
	"fmt"
	"strings"
)

func Mostrepeatedchar() {
	name := "banan"
	name = strings.ReplaceAll(name, " ", "")
	frequency := make(map[rune]int)
	maxcount := 0
	for _, char := range name {
		frequency[char]++
		if maxcount < frequency[char] {
			maxcount = frequency[char]
		}
	}
	var slice []rune
	for char, count := range frequency {
		if count == maxcount {
			slice = append(slice, char)
		}
	}
	fmt.Printf("%c ", slice)
}
