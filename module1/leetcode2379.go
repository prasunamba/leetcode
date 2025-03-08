package module1

import (
	"fmt"
	"math"
)

func Leetcode2379() {
	block := "BWWWBB"
	k := 6

	min := math.MaxInt
	blockslice := []rune(block)
	for i := 0; i <= len(blockslice)-k; i++ {
		white := 0
		for _, val := range blockslice[i : i+k] {
			if val == 'W' {
				white++
			}
		}
		if white < min {
			min = white
		}
	}
	fmt.Println("", min)
}
