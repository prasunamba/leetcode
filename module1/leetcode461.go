package module1

import (
	"fmt"
	"strconv"
)

func HammingDistance(x int, y int) int {
	a := int64(x)
	binarya := fmt.Sprintf("%b", a)
	binaryb := strconv.FormatInt(int64(y), 2)
	fmt.Println("b", binaryb)
	maxLen := len(binarya) // Get the longest binary length
	if len(binaryb) > maxLen {
		maxLen = len(binaryb)
	}

	// Format to pad with leading zeros
	binA := fmt.Sprintf("%0*s", maxLen, binarya)
	binB := fmt.Sprintf("%0*s", maxLen, binaryb)
	fmt.Println("bin", binA, binB)
	slicea := []rune(binA)
	sliceb := []rune(binB)

	hammdist := 0
	for idx, val := range slicea {
		if val != sliceb[idx] {
			hammdist++
		}
	}
	return hammdist
}
