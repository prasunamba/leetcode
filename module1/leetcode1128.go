package module1

import (
	"fmt"
	"sort"
)

// my version not optimized
func NumEquivDominoPairs(dominoes [][]int) int {
	seenmp := make(map[int]bool)
	freqmap := make(map[int]int)
	for _, list := range dominoes {
		sort.Ints(list)
		key := list[0]*10 + list[1]
		seenmp[key] = true
		freqmap[key]++
	}
	result := []int{}

	for _, value := range freqmap {
		fmt.Println("", value, result)
		if value > 1 {
			result = append(result, value)
		}
	}
	op := 0
	for _, value := range result {

		op += (value * (value - 1)) / 2
	}
	return op
}

// optimized version
/* func numEquivDominoPairs(dominoes [][]int) int {
	count := make(map[int]int)
	res := 0

	for _, d := range dominoes {
		a, b := d[0], d[1]
		// Always store with smaller number first to handle [a,b] == [b,a]
		key := min(a, b)*10 + max(a, b)
		res += count[key]
		fmt.Println("res", res, count)
		count[key]++
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
} */
